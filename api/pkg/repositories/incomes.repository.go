package repository

import (
	"fmt"
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

func (r *incomeRepository) Update(id uuid.UUID, updates interface{}) error {
	result := r.db.Model(&models.Income{}).Where("id = ?", id).Updates(updates)
	fmt.Println(result.Error, updates)
	return result.Error
}

func (r *incomeRepository) Delete(id uuid.UUID) error {
	result := r.db.Where("id = ?", id).Delete(&models.Income{})
	return result.Error
}
