package helpers

import (
	"sort"

	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
)

// GetEmployeeToReportsLevelMap arranges a list of ManagerChains into a dictionary mapping
// each manager to their reports, excluding self-relations (level 0).
//
// Given a list of ManagerChains, this function creates a dictionary where the keys are
// ManagerIds and the values are slices of TupleProcessors1 containing EmployeeId and Level.
// ManagerChains with level 0 are excluded from the result.
//
// Parameters:
// - manager_chains: a slice of ManagerChain instances representing the manager chains.
//
// Returns:
//   - map[int][]security_model.TupleProcessors1: a map where the keys are ManagerIds (int)
//     and the values are slices of TupleProcessors1 containing EmployeeId (uint) and Level (int).
//   - error: an error if there was any issue processing the input.
func GetEmployeeToReportsLevelMap(manager_chains []model.ManagerChain) (map[int][]security_model.TupleProcessors1, error) {
	ret := make(map[int][]security_model.TupleProcessors1)

	for _, manager_chain := range manager_chains {
		if manager_chain.Level == 0 {
			continue
		}

		key := manager_chain.ManagerId
		ret[key] = append(ret[key], security_model.TupleProcessors1{
			EmployeeId: uint(manager_chain.EmployeeId),
			Level:      manager_chain.Level,
		})
	}

	for _, tuples := range ret {
		sort.Slice(tuples, func(i, j int) bool {
			return tuples[i].EmployeeId < tuples[j].EmployeeId
		})
	}

	return ret, nil
}

// GetEmployeeToManagersLevelMap arranges a list of ManagerChains into a dictionary mapping
// each employee to their managers, excluding self-relations (level 0).
//
// Given a list of ManagerChains, this function creates a dictionary where the keys are
// EmployeeIds and the values are slices of TupleProcessors2 containing ManagerId and Level.
// ManagerChains with level 0 are excluded from the result.
//
// Parameters:
// - manager_chains: a slice of ManagerChain instances representing the manager chains.
//
// Returns:
//   - map[int][]security_model.TupleProcessors2: a map where the keys are EmployeeIds (int)
//     and the values are slices of TupleProcessors2 containing ManagerId (uint) and Level (int).
//   - error: an error if there was any issue processing the input.
func GetEmployeeToManagersLevelMap(manager_chains []model.ManagerChain) (map[int][]security_model.TupleProcessors2, error) {
	ret := make(map[int][]security_model.TupleProcessors2)

	for _, manager_chain := range manager_chains {
		if manager_chain.Level == 0 {
			continue
		}
		key := manager_chain.EmployeeId
		ret[key] = append(ret[key], security_model.TupleProcessors2{
			ManagerId: uint(manager_chain.ManagerId),
			Level:     manager_chain.Level,
		})
	}

	return ret, nil
}
