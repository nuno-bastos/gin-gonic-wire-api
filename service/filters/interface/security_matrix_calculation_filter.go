package filters_interface

import (
	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
)

type SecurityMatrixCalculationFilter interface {
	// ExecuteFilter applies the filter logic to the given profile, user, and input data, creating rule(s) to determine access control.
	//
	// Parameters:
	// - profile: a pointer to the Profile model instance to be filtered.
	// - user: a pointer to the User model instance to be filtered.
	// - input: an instance of SecurityMatrixCalculationFilterInput.
	//
	// Returns:
	// - ProfileUserAllowOrDenyRules: the result of the filter operation, indicating whether the user is allowed or denied access.
	ExecuteFilter(profile *model.Profile, user *model.User, input security_model.SecurityMatrixCalculationFilterInput) security_model.ProfileUserAllowOrDenyRules
}
