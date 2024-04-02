package interfaces

import (
	"gofinance/api/pkg/models"
	"gofinance/api/pkg/types"

	"github.com/google/uuid"
)

type ExpenseRepository interface {
	Create(*models.Expense) (*models.Expense, error)
	FindByID(uuid.UUID) (*models.Expense, error)
	Update(*models.Expense) error
	Delete(id uuid.UUID) error
}

type ExpenseServices interface {
	Create(*types.ExpenseDTO) (*types.ExpenseDTO, error)
	FindByID(uuid.UUID) (*types.ExpenseDTO, error)
	Update(*types.ExpenseDTO, uuid.UUID) (*types.ExpenseDTO, error)
	Delete(id uuid.UUID) error
}
