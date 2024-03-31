package repository

import (
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) interfaces.ExpenseRepository {
	return &expenseRepository{db}
}

func (r *expenseRepository) Create(expense *models.Expense) (*models.Expense, error) {
	result := r.db.Create(expense)
	if result.Error != nil {
		return nil, result.Error
	}
	return expense, nil
}

func (r *expenseRepository) FindByID(id uuid.UUID) (*models.Expense, error) {
	var expense models.Expense
	err := r.db.First(&expense, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *expenseRepository) FindAll() ([]*models.Expense, error) {
	var expenses []*models.Expense
	err := r.db.Find(&expenses).Error
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (r *expenseRepository) Update(id uuid.UUID, updates interface{}) error {
	result := r.db.Where("id = ?", id).Updates(updates)
	return result.Error
}