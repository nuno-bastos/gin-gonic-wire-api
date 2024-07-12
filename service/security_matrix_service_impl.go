package service

import (
	"context"
	"log"
	"runtime"
	"sync"

	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
	interfaces "github.com/nuno-bastos/gin-gonic-wire-api/repo/interface"
	filters_interface "github.com/nuno-bastos/gin-gonic-wire-api/service/filters/interface"
	helpers "github.com/nuno-bastos/gin-gonic-wire-api/service/helpers"
	services "github.com/nuno-bastos/gin-gonic-wire-api/service/interface"
)

type securityMatrixCalculatorService struct {
	managerChainRepo interfaces.ManagerChainRepository
	userRepo         interfaces.UserRepository
	ruleRepo         interfaces.SecurityMatrixRuleRepository
	managerFilters   []filters_interface.IManagerFilter
	employeeFilters  []filters_interface.IEmployeeFilter
}

// ProfileTypeId defines the IDs for different types of profiles.var ProfileTypeId = struct {
var ProfileTypeId = struct {
	ADMIN    string
	MANAGER  string
	EMPLOYEE string
}{
	ADMIN:    "PFT11111111111111111",
	MANAGER:  "PFT22222222222222222",
	EMPLOYEE: "PFT33333333333333333",
}

// NewSecurityMatrixCalculatorService creates a new instance of SecurityMatrixCalculatorService.
//
// Parameters:
// - _managerChainRepo: The ManagerChainRepository.
// - _userRepo: The UserRepository.
// - _ruleRepo: The SecurityMatrixRuleRepository.
// - _managerFilters: The list of manager filters.
// - _employeeFilters: The list of employee filters.
//
// Returns:
// - services.SecurityMatrixService: The new instance of SecurityMatrixService.
func NewSecurityMatrixCalculatorService(
	_managerChainRepo interfaces.ManagerChainRepository,
	_userRepo interfaces.UserRepository,
	_ruleRepo interfaces.SecurityMatrixRuleRepository,
	_managerFilters []filters_interface.IManagerFilter,
	_employeeFilters []filters_interface.IEmployeeFilter,
) services.SecurityMatrixService {
	return &securityMatrixCalculatorService{
		managerChainRepo: _managerChainRepo,
		userRepo:         _userRepo,
		ruleRepo:         _ruleRepo,
		managerFilters:   _managerFilters,
		employeeFilters:  _employeeFilters,
	}
}

// CalculateGoSecurityMatrix calculates and writes the security matrix rules to the database.
// Replicates the original service Handle function.
//
// Parameters:
// - ctx: The context for the operation.
func (p *securityMatrixCalculatorService) CalculateGoSecurityMatrix(ctx context.Context) {
	// Fetch manager chains from repository
	manager_chains, err := p.managerChainRepo.FindAll(ctx)
	if err != nil {
		log.Fatalf("Error fetching manager chains: %v", err)
	}

	// Fetch users who are employees and have associated profiles
	employees_with_profiles, err := p.userRepo.FetchEmployeesWithAssociatedProfiles(ctx)
	if err != nil {
		log.Fatalf("Error fetching users with associated profiles: %v", err)
	}

	// a List of Tuple<Profile, User> where the profile is of type Manager
	var managers_profiles_users []*security_model.TupleProfileUser
	// a List of Tuple<Profile, User> where the profile is of type Employee (there should be only one such profile)
	var employees_profiles_users []*security_model.TupleProfileUser

	for _, user := range employees_with_profiles {
		for _, profile := range user.Profiles {
			if profile.ProfileTypeId == ProfileTypeId.MANAGER {
				managers_profiles_users = append(managers_profiles_users, &security_model.TupleProfileUser{
					Profile: profile,
					User:    user,
				})
			} else if profile.ProfileTypeId == ProfileTypeId.EMPLOYEE {
				employees_profiles_users = append(employees_profiles_users, &security_model.TupleProfileUser{
					Profile: profile,
					User:    user,
				})
			}
		}
	}

	// Security Matrix Calculation Input
	// 3 Dictionaries derived from managerChains
	managers_map, err := helpers.GetEmployeeToManagersLevelMap(manager_chains)
	if err != nil {
		log.Fatalf("Error generating managers map: %v", err)
	}
	reports_map, err := helpers.GetEmployeeToReportsLevelMap(manager_chains)
	if err != nil {
		log.Fatalf("Error generating reports map: %v", err)
	}
	levels_map, err := helpers.GetLevelToEmployeesMap(employees_with_profiles)
	if err != nil {
		log.Fatalf("Error generating level to employees map: %v", err)
	}

	input := security_model.SecurityMatrixCalculationFilterInput{
		EmployeeToManagers: managers_map,
		EmployeeToReports:  reports_map,
		LevelToEmployees:   levels_map,
	}

	rules_list := p.CalculateRules(managers_profiles_users, employees_profiles_users, input)
	flattened_rules := helpers.Flatten(rules_list)
	final_rules := RemoveRulesWithDefaultValues(flattened_rules)
	if err := p.ruleRepo.WriteRules(ctx, final_rules); err != nil {
		log.Fatalf("Error writing rules to database: %v", err)
	}
}

// CalculateRules calculates security matrix rules for managers and employees by applying filters.
//
// Parameters:
// - managersProfilesUsers: List of manager profiles and users.
// - employeesProfilesUsers: List of employee profiles and users.
// - input: The SecurityMatrixCalculationFilterInput.
//
// Returns:
// - []*security_model.ProfileUserAllowOrDenyRules: List of calculated rules.
//
// Function Logic:
// It concurrently calculates security matrix rules for managers and employees using filters.
// It creates worker goroutines to process the input data concurrently. The number of workers is determined by the number of available CPU cores.
// Each worker receives a channel of TupleProfileUser and sends calculated rules to the rules channel.
// The profile user data is distributed to workers through a channel.
// Once all data is processed, the rules channel is closed and the collected rules are returned.
func (p *securityMatrixCalculatorService) CalculateRules(
	managersProfilesUsers, employeesProfilesUsers []*security_model.TupleProfileUser,
	input security_model.SecurityMatrixCalculationFilterInput,
) []*security_model.ProfileUserAllowOrDenyRules {
	// Declare a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Create a channel to receive calculated rules
	rules_channel := make(chan *security_model.ProfileUserAllowOrDenyRules, len(managersProfilesUsers)+len(employeesProfilesUsers))

	// Create a channel to distribute profile user data to workers
	profile_user_channel := make(chan *security_model.TupleProfileUser, len(managersProfilesUsers)+len(employeesProfilesUsers))

	// Determine the number of worker goroutines based on available CPU cores
	num_workers := runtime.NumCPU()

	// Define the worker function
	worker := func(profile_user_channel <-chan *security_model.TupleProfileUser, rules_channel chan<- *security_model.ProfileUserAllowOrDenyRules) {
		defer wg.Done() // Notify the wait group when the worker finishes

		// Process each profile user received from the channel
		for profile_user := range profile_user_channel {
			// Execute filters based on the profile type and send the calculated rules to the rules channel
			if profile_user.Profile.ProfileTypeId == ProfileTypeId.MANAGER {
				for _, filter := range p.managerFilters {
					result := filter.ExecuteFilter(profile_user.Profile, profile_user.User, input)
					rules_channel <- &result
				}
			} else if profile_user.Profile.ProfileTypeId == ProfileTypeId.EMPLOYEE {
				for _, filter := range p.employeeFilters {
					result := filter.ExecuteFilter(profile_user.Profile, profile_user.User, input)
					rules_channel <- &result
				}
			}
		}
	}

	// Create worker goroutines
	for i := 0; i < num_workers; i++ {
		wg.Add(1) // Increment the wait group counter for each worker
		go worker(profile_user_channel, rules_channel)
	}

	// Distribute profile_user data to workers
	go func() {
		// Send profile user data to the profile user channel
		for _, profile_user := range append(managersProfilesUsers, employeesProfilesUsers...) {
			profile_user_channel <- profile_user
		}
		close(profile_user_channel) // Close the channel after sending all data
	}()

	// Collect calculated rules from the rules channel
	var rules_list []*security_model.ProfileUserAllowOrDenyRules
	go func() {
		wg.Wait()            // Wait for all workers to finish
		close(rules_channel) // Close the rules channel after all rules are calculated
	}()

	// Receive rules from the channel and append them to the rules list
	for result := range rules_channel {
		rules_list = append(rules_list, result)
	}

	// Return the collected rules
	return rules_list
}

// RemoveRulesWithDefaultValues removes rules with default access level values.
//
// Parameters:
// - param_rules: List of rules to filter.
//
// Returns:
// - []security_model.GoCorsicaMatrixRule: Filtered list of rules.
func RemoveRulesWithDefaultValues(param_rules []security_model.GoMatrixRule) []security_model.GoMatrixRule {
	var rules_list []security_model.GoMatrixRule
	for _, rule := range param_rules {
		if rule.AccessLevelRead != uint32(0) {
			rules_list = append(rules_list, rule)
		}
	}
	return rules_list
}
