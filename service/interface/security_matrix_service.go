package interfaces

import (
	"context"

	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
)

type SecurityMatrixService interface {
	CalculateGoSecurityMatrix(ctx context.Context)
	CalculateRules(managersProfilesUsers, employeesProfilesUsers []*security_model.TupleProfileUser, input security_model.SecurityMatrixCalculationFilterInput) []*security_model.ProfileUserAllowOrDenyRules
}
