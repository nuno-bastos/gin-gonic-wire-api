package filters

import (
	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
	filters_interface "github.com/nuno-bastos/gin-gonic-wire-api/service/filters/interface"
)

type DenyManagersFilter struct{}

func NewDenyManagersFilter() *DenyManagersFilter {
	return &DenyManagersFilter{}
}

// ExecuteFilter generates deny rules for the given profile and user to deny access
// to other managers who manage the same employee. The rules prevent the manager
// from accessing other managers who manage the same employee.
//
// Parameters:
// - profile: a pointer to the Profile model instance.
// - user: a pointer to the User model instance representing the manager.
// - input: an instance of SecurityMatrixCalculationFilterInput containing filter criteria.
//
// Returns:
//   - ProfileUserAllowOrDenyRules: the result of the filter operation, containing the generated
//     deny rules for the manager to deny access to other managers who manage the same employee.
func (f *DenyManagersFilter) ExecuteFilter(profile *model.Profile, user *model.User, input security_model.SecurityMatrixCalculationFilterInput) security_model.ProfileUserAllowOrDenyRules {
	var rules []security_model.AllowOrDenyRule // Initialize a list to store the generated rules
	employeeId := int(user.EmployeeId)         // Retrieve the employee ID associated with the userEmployee

	// Check if the EmployeeToManagersLevelDictionary contains the employeeId
	if managers, found := input.EmployeeToManagers[employeeId]; found {
		// Iterate over each manager associated with the employee
		for _, manager := range managers {
			rules = append(rules, security_model.AllowOrDenyRule{
				UserManagerId:   user.Id,
				EmployeeId:      manager.ManagerId,
				AccessLevelRead: security_model.DENY_ACCESS_LEVEL_READ,
			})
		}
	}

	return security_model.ProfileUserAllowOrDenyRules{
		ProfileId:        profile.Id,
		UserId:           user.Id,
		OriginFilter:     "DenyManagersFilter",
		AllowOrDenyRules: rules,
	}
}

// Ensure DenyManagersFilter implements the IManagerFilter interface.
var _ filters_interface.IManagerFilter = (*DenyManagersFilter)(nil)
