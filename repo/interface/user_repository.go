package interfaces

import (
	"context"

	"github.com/nuno-bastos/gin-gonic-wire-api/model"
)

type UserRepository interface {
	FetchEmployeesWithAssociatedProfiles(ctx context.Context) ([]*model.User, error)
}
