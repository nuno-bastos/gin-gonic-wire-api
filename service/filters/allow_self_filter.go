package filters

import (
	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
	filters_interface "github.com/nuno-bastos/gin-gonic-wire-api/service/filters/interface"
)

type AllowSelfFilter struct{}

func NewAllowSelfFilter() *AllowSelfFilter {
	return &AllowSelfFilter{}
}

// ExecuteFilter generates allow rules for the given profile and user to access
// the user's own profile. The rule allows the user full access to their own profile.
//
// Parameters:
// - profile: a pointer to the Profile model instance.
// - user: a pointer to the User model instance representing the employee.
// - input: an instance of SecurityMatrixCalculationFilterInput containing filter criteria.
//
// Returns:
//   - ProfileUserAllowOrDenyRules: the result of the filter operation, containing the generated
//     allow rule for the employee to access their own profile.
func (f *AllowSelfFilter) ExecuteFilter(profile *model.Profile, user *model.User, input security_model.SecurityMatrixCalculationFilterInput) security_model.ProfileUserAllowOrDenyRules {
	rules := []security_model.AllowOrDenyRule{
		{
			UserManagerId:   user.Id,
			EmployeeId:      user.EmployeeId,
			AccessLevelRead: uint32(security_model.All), // Full read access
		},
	}

	return security_model.ProfileUserAllowOrDenyRules{
		ProfileId:        profile.Id,
		UserId:           user.Id,
		OriginFilter:     "AllowSelfFilter",
		AllowOrDenyRules: rules,
	}
}

// Ensure AllowSelfFilter implements the IEmployeeFilter interface.
var _ filters_interface.IEmployeeFilter = (*AllowSelfFilter)(nil)
