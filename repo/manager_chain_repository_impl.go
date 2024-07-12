package repo

import (
	"context"

	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	interfaces "github.com/nuno-bastos/gin-gonic-wire-api/repo/interface"

	"gorm.io/gorm"
)

// managerChainRepository provides methods for accessing and manipulating ManagerChains
type managerChainRepository struct {
	DB *gorm.DB
}

func NewManagerChainRepository(DB *gorm.DB) interfaces.ManagerChainRepository {
	return &managerChainRepository{DB}
}

// FindAll retrieves all ManagerChain records from the database.
// Parameters:
// - ctx: context for managing request lifecycle.
//
// Returns:
// - []model.ManagerChain: a slice of ManagerChain records.
// - error: error object if the operation fails, nil otherwise.
func (p *managerChainRepository) FindAll(ctx context.Context) ([]model.ManagerChain, error) {
	var managerChains []model.ManagerChain
	err := p.DB.Find(&managerChains).Error

	return managerChains, err
}
