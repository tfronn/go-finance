package services

import (
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/models"
	"gofinance/api/pkg/types"

	"github.com/google/uuid"
)

type expenseServices struct {
	expenseRepo interfaces.ExpenseRepository
}

func NewExpenseService(repo interfaces.ExpenseRepository) interfaces.ExpenseServices {
	return &expenseServices{expenseRepo: repo}
}

func (ps *expenseServices) Create(expense *types.ExpenseDTO) (*types.ExpenseDTO, error) {

	newExpense := &models.Expense{
		ID:        expense.ID,
		Amount:    expense.Amount,
		Category:  expense.Category,
		CreatedAt: expense.CreatedAt,
		UserID:    expense.UserID,
	}

	_, err := ps.expenseRepo.Create(newExpense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (ps *expenseServices) FindByID(id uuid.UUID) (*types.ExpenseDTO, error) {
	result, err := ps.expenseRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	foundExpense := &types.ExpenseDTO{
		ID:        result.ID,
		Amount:    result.Amount,
		Category:  result.Category,
		CreatedAt: result.CreatedAt,
	}
	return foundExpense, nil
}

func (ps *expenseServices) Update(expense *types.ExpenseDTO, updates interface{}) (*types.ExpenseDTO, error) {
	err := ps.expenseRepo.Update(expense.ID, updates)
	if err != nil {
		return expense, err
	}

	result, err := ps.expenseRepo.FindByID(expense.ID)
	if err != nil {
		return nil, err
	}

	foundExpense := &types.ExpenseDTO{
		ID:        result.ID,
		Amount:    result.Amount,
		Category:  result.Category,
		CreatedAt: result.CreatedAt,
		UserID:    result.UserID,
	}

	return foundExpense, nil
}

func (ps *expenseServices) Delete(id uuid.UUID) error {
	err := ps.expenseRepo.Delete(id)
	return err
}
