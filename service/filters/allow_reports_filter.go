package filters

import (
	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
	filters_interface "github.com/nuno-bastos/gin-gonic-wire-api/service/filters/interface"
)

type AllowReportsFilter struct{}

func NewAllowReportsFilter() *AllowReportsFilter {
	return &AllowReportsFilter{}
}

// ExecuteFilter generates allow rules for the given profile and user based on the reports
// associated with the user in the input data. The rules allow the manager to access the
// reports or profiles of their subordinate users.
//
// Parameters:
// - profile: a pointer to the Profile model instance.
// - user: a pointer to the User model instance representing the manager.
// - input: an instance of SecurityMatrixCalculationFilterInput containing filter criteria.
//
// Returns:
//   - ProfileUserAllowOrDenyRules: the result of the filter operation, containing the generated
//     allow rules for the manager to access reports or profiles of their subordinate users.
func (f *AllowReportsFilter) ExecuteFilter(profile *model.Profile, user *model.User, input security_model.SecurityMatrixCalculationFilterInput) security_model.ProfileUserAllowOrDenyRules {
	var rules []security_model.AllowOrDenyRule // Initialize a list to store the generated rules
	employeeId := user.EmployeeId              // Retrieve the employee ID from the user

	// Check if the employee has reports in the input dictionary
	if reports, found := input.EmployeeToReports[int(employeeId)]; found {
		// Iterate over each report in the list of reports
		for _, report := range reports {
			// Create an AllowRule for the user and the report, using the access levels from the profile
			rules = append(rules, security_model.AllowOrDenyRule{
				UserManagerId:   user.Id,
				EmployeeId:      report.EmployeeId,
				AccessLevelRead: uint32(profile.AccessLevelRead),
			})
		}
	}

	return security_model.ProfileUserAllowOrDenyRules{
		ProfileId:        profile.Id,
		UserId:           user.Id,
		OriginFilter:     "AllowReportsFilter",
		AllowOrDenyRules: rules,
	}
}

// Ensure AllowReportsFilter implements the IManagerFilter interface.
var _ filters_interface.IManagerFilter = (*AllowReportsFilter)(nil)
