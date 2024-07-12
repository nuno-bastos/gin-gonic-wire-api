package helpers

import model "github.com/nuno-bastos/gin-gonic-wire-api/model"

// GetLevelToEmployeesMap generates a map where each profile level is a key
// and the corresponding employee IDs are stored as values.
//
// Given a list of users, each potentially having multiple profiles with different levels,
// this function returns a map with each profile level as a key and a list of employee IDs
// that have profiles with that level.
//
// Parameters:
// - employees_with_profiles: a slice of User pointers representing employees, each potentially
//   having multiple profiles with different levels.
//
// Returns:
// - map[uint][]uint: a map where the keys are profile levels (uint) and the values are slices of
//   employee IDs (uint) associated with those levels.
// - error: an error if there was any issue processing the input.
func GetLevelToEmployeesMap(employees_with_profiles []*model.User) (map[uint][]uint, error) {
	result_map := make(map[uint][]uint)

	for _, user := range employees_with_profiles {
		for _, profile := range user.Profiles {
			result_map[profile.Level] = append(result_map[profile.Level], user.EmployeeId)
		}
	}

	return result_map, nil
}
