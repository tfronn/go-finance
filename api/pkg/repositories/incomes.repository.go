package repository

import (
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type incomeRepository struct {
	db *gorm.DB
}

func NewIncomeRepository(db *gorm.DB) interfaces.IncomeRepository {
	return &incomeRepository{db}
}

func (r *incomeRepository) Create(income *models.Income) (*models.Income, error) {
	result := r.db.Create(income)
	if result.Error != nil {
		return nil, result.Error
	}
	return income, nil
}

func (r *incomeRepository) FindByID(id uuid.UUID) (*models.Income, error) {
	var income models.Income
	err := r.db.First(&income, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &income, nil
}

func (r *incomeRepository) FindAll() ([]*models.Income, error) {
	var incomes []*models.Income
	err := r.db.Find(&incomes).Error
	if err != nil {
		return nil, err
	}
	return incomes, nil
}

func (r *incomeRepository) Update(id uuid.UUID, updates interface{}) error {
	result := r.db.Where("id = ?", id).Updates(updates)
	return result.Error
}
