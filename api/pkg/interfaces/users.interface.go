package interfaces

import (
	"gofinance/api/pkg/models"
	"gofinance/api/pkg/types"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(*models.User) (*models.User, error)
	FindByID(uuid.UUID) (*models.User, error)
	FindAll() ([]*models.User, error)
	FindByGoogleID(googleID string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(uuid.UUID, interface{}) error
}

type UserServices interface {
	Create(*types.UserDTO) (*types.UserDTO, error)
	FindByID(uuid.UUID) (*types.UserDTO, error)
	FindAll() ([]*types.UserDTO, error)
	FindByGoogleID(googleID string) (*types.UserDTO, error)
	FindByEmail(email string) (*types.UserDTO, error)
	FindByToken(tokenString string) (*types.UserDTO, error)
	Update(*types.UserDTO, interface{}) (*types.UserDTO, error)
}
