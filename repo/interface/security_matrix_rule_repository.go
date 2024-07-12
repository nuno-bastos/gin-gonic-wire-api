package interfaces

import (
	"context"

	model_security "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
)

type SecurityMatrixRuleRepository interface {
	Migrate() error
	WriteRules(ctx context.Context, rules []model_security.GoMatrixRule) error
}
