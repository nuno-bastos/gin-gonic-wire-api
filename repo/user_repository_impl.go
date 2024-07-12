package repo

import (
	"context"

	model "github.com/nuno-bastos/gin-gonic-wire-api/model"
	interfaces "github.com/nuno-bastos/gin-gonic-wire-api/repo/interface"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{DB}
}

// FetchEmployeesWithAssociatedProfiles retrieves all Users who have associated Profiles.
// Parameters:
// - ctx: context for managing request lifecycle.
//
// Returns:
// - []*model.User: a slice of User records with their associated Profiles.
// - error: error object if the operation fails, nil otherwise.
func (p *userRepository) FetchEmployeesWithAssociatedProfiles(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	// Setup the join table for the many-to-many relationship between Users and Profiles
	p.DB.SetupJoinTable(&model.User{}, "Profiles", &model.ProfileUser{})

	// Retrieve users with associated profiles where EmployeeId is not null and greater than 0
	result := p.DB.Preload("Profiles").Where("EmployeeId IS NOT NULL AND EmployeeId > 0").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
