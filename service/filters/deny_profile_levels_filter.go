package filters

import (
	"sort"

	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
	filters_interface "github.com/nuno-bastos/gin-gonic-wire-api/service/filters/interface"
)

type DenyProfileLevelsFilter struct{}

func NewDenyProfileLevelsFilter() *DenyProfileLevelsFilter {
	return &DenyProfileLevelsFilter{}
}

// ExecuteFilter generates deny rules for the given profile and user to deny access
// to all other employees who have a lower or equal access level. The rules prevent
// the manager from accessing profiles of other employees who are at the same or lower
// access level.
//
// Parameters:
// - profile: a pointer to the Profile model instance.
// - user: a pointer to the User model instance representing the manager.
// - input: an instance of SecurityMatrixCalculationFilterInput containing filter criteria.
//
// Returns:
//   - ProfileUserAllowOrDenyRules: the result of the filter operation, containing the generated
//     deny rules for the manager to deny access to other employees' profiles.
func (f *DenyProfileLevelsFilter) ExecuteFilter(profile *model.Profile, user *model.User, input security_model.SecurityMatrixCalculationFilterInput) security_model.ProfileUserAllowOrDenyRules {
	var rules []security_model.AllowOrDenyRule

	// Sort the levels for iteration in ascending order
	var levels []uint
	for level := range input.LevelToEmployees {
		levels = append(levels, level)
	}
	sort.Slice(levels, func(i, j int) bool {
		return levels[i] < levels[j]
	})

	// Iterate over the sorted levels and create deny rules
	for _, level := range levels {
		if level <= uint(profile.Level) {
			for _, employeeId := range input.LevelToEmployees[level] {
				rules = append(rules, security_model.AllowOrDenyRule{
					UserManagerId:   user.Id,
					EmployeeId:      employeeId,
					AccessLevelRead: security_model.DENY_ACCESS_LEVEL_READ,
				})
			}
		}
	}

	return security_model.ProfileUserAllowOrDenyRules{
		ProfileId:        profile.Id,
		UserId:           user.Id,
		OriginFilter:     "DenyProfileLevelsFilter",
		AllowOrDenyRules: rules,
	}
}

// Ensure DenyProfileLevelsFilter implements the IManagerFilter interface.
var _ filters_interface.IManagerFilter = (*DenyProfileLevelsFilter)(nil)
