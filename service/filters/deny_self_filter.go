package filters

import (
	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
	filters_interface "github.com/nuno-bastos/gin-gonic-wire-api/service/filters/interface"
)

type DenySelfFilter struct{}

func NewDenySelfFilter() *DenySelfFilter {
	return &DenySelfFilter{}
}

// ExecuteFilter generates deny rules for the given profile and user to deny access
// to the user's own profile. The rule prevents the employee from accessing their own profile.
//
// Parameters:
// - profile: a pointer to the Profile model instance.
// - user: a pointer to the User model instance representing the employee.
// - input: an instance of SecurityMatrixCalculationFilterInput containing filter criteria.
//
// Returns:
//   - ProfileUserAllowOrDenyRules: the result of the filter operation, containing the generated
//     deny rule for the employee to deny access to their own profile.
func (f *DenySelfFilter) ExecuteFilter(profile *model.Profile, user *model.User, input security_model.SecurityMatrixCalculationFilterInput) security_model.ProfileUserAllowOrDenyRules {
	rules := []security_model.AllowOrDenyRule{
		{
			UserManagerId:   user.Id,
			EmployeeId:      user.EmployeeId,
			AccessLevelRead: uint32(security_model.None),
		},
	}

	return security_model.ProfileUserAllowOrDenyRules{
		ProfileId:        profile.Id,
		UserId:           user.Id,
		OriginFilter:     "DenySelfFilter",
		AllowOrDenyRules: rules,
	}
}

// Ensure DenySelfFilter implements the IEmployeeFilter interface.
var _ filters_interface.IEmployeeFilter = (*DenySelfFilter)(nil)
