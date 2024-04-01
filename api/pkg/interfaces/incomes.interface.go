package interfaces

import (
	"gofinance/api/pkg/models"
	"gofinance/api/pkg/types"

	"github.com/google/uuid"
)

type IncomeRepository interface {
	Create(*models.Income) (*models.Income, error)
	FindByID(uuid.UUID) (*models.Income, error)
	Update(uuid.UUID, interface{}) error
	Delete(id uuid.UUID) error
}

type IncomeServices interface {
	Create(*types.IncomeDTO) (*types.IncomeDTO, error)
	FindByID(uuid.UUID) (*types.IncomeDTO, error)
	Update(*types.IncomeDTO, interface{}) (*types.IncomeDTO, error)
	Delete(id uuid.UUID) error
}
