package interfaces

import (
	"context"

	"github.com/nuno-bastos/gin-gonic-wire-api/model"
)

type ManagerChainRepository interface {
	FindAll(ctx context.Context) ([]model.ManagerChain, error)
}
