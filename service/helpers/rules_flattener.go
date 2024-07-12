package helpers

import (
	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
)

// Flatten takes a collection of ProfileUserAllowOrDenyRules, groups them by UserManagerId and EmployeeId,
// and then calculates the flattened MatrixRules based on the access levels defined in those rules.
//
// Parameters:
// - input_rules: A slice of ProfileUserAllowOrDenyRules containing rules to be flattened.
//
// Returns:
// - []security_model.GoMatrixRule: A slice of GoMatrixRule containing the flattened rules.
func Flatten(input_rules []*security_model.ProfileUserAllowOrDenyRules) []security_model.GoMatrixRule {
	// A map to store the condensed rules
	grouped_rules := make(map[string]map[uint][]security_model.AllowOrDenyRule)

	// Iterate over the input rules
	for _, profile_rules := range input_rules {
		for _, rule := range profile_rules.AllowOrDenyRules {
			if _, ok := grouped_rules[rule.UserManagerId]; !ok {
				grouped_rules[rule.UserManagerId] = make(map[uint][]security_model.AllowOrDenyRule)
			}

			grouped_rules[rule.UserManagerId][rule.EmployeeId] = append(grouped_rules[rule.UserManagerId][rule.EmployeeId], rule)
		}
	}

	// Convert the condensed rules into MatrixRules
	var flattened_rules []security_model.GoMatrixRule
	for userManagerId, employeeMap := range grouped_rules {
		for employeeId, rules := range employeeMap {
			accessLevel := CalculateReadAccessLevel(rules)

			flattened_rules = append(flattened_rules, security_model.GoMatrixRule{
				UserId:          userManagerId,
				EmployeeId:      employeeId,
				AccessLevelRead: accessLevel,
			})
		}
	}

	return flattened_rules
}

// CalculateReadAccessLevel calculates the read access level based on the input rules.
//
// Parameters:
// - rules: A slice of AllowOrDenyRule to calculate the access level from.
//
// Returns:
// - uint32: The calculated read access level.
//
// Function Logic:
// It iterates over the input rules, condensing them into allow and deny access levels using bitwise operations.
// It starts with FULL_DENY as the initial access level and applies bitwise AND and NOT operations to calculate
// the final access level.
func CalculateReadAccessLevel(rules []security_model.AllowOrDenyRule) uint32 {
	condensed_allow, condensed_deny := Condense(rules)

	accessLevel := security_model.FULL_DENY
	accessLevel &= condensed_allow
	accessLevel &= ^condensed_deny

	return accessLevel
}

// Condense condenses the input rules into allow and deny access levels using bitwise operations.
//
// Parameters:
// - rules: A slice of AllowOrDenyRule to condense.
//
// Returns:
// - uint32: The condensed allow access level.
// - uint32: The condensed deny access level.
//
// Function Logic:
// It iterates over the input rules, applying bitwise OR operations to condense the allow and deny access levels.
// If a rule's is of deny type (DENY_ACCESS_LEVEL_READ), it applies a bitwise OR to the deny access level;
// otherwise, it applies a bitwise OR to the allow access level.
func Condense(rules []security_model.AllowOrDenyRule) (allow uint32, deny uint32) {
	for _, rule := range rules {
		if rule.AccessLevelRead == security_model.FULL_DENY {
			deny = security_model.FULL_DENY
			continue
		}

		if rule.AccessLevelRead == security_model.DENY_ACCESS_LEVEL_READ {
			deny |= rule.AccessLevelRead
		} else {
			allow |= rule.AccessLevelRead
		}
	}
	return allow, deny
}
