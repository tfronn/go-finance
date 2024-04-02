package services

import (
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/models"
	"gofinance/api/pkg/types"

	"github.com/google/uuid"
)

type incomeServices struct {
	incomeRepo interfaces.IncomeRepository
}

func NewIncomeService(repo interfaces.IncomeRepository) interfaces.IncomeServices {
	return &incomeServices{incomeRepo: repo}
}

func (ps *incomeServices) Create(income *types.IncomeDTO) (*types.IncomeDTO, error) {

	newIncome := &models.Income{
		ID:          income.ID,
		Description: income.Description,
		Amount:      income.Amount,
		Category:    income.Category,
		CreatedAt:   income.CreatedAt,
		UserID:      income.UserID,
	}

	_, err := ps.incomeRepo.Create(newIncome)
	if err != nil {
		return nil, err
	}

	return income, nil
}

func (ps *incomeServices) FindByID(id uuid.UUID) (*types.IncomeDTO, error) {
	result, err := ps.incomeRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	foundIncome := &types.IncomeDTO{
		ID:          result.ID,
		Description: result.Description,
		Amount:      result.Amount,
		Category:    result.Category,
		CreatedAt:   result.CreatedAt,
		UserID:      result.UserID,
	}
	return foundIncome, nil
}

func (ps *incomeServices) Update(income *types.IncomeDTO, id uuid.UUID) (*types.IncomeDTO, error) {
	incomeToUpdate := &models.Income{
		ID:          income.ID,
		Description: income.Description,
		Amount:      income.Amount,
		Category:    income.Category,
		CreatedAt:   income.CreatedAt,
		UserID:      income.UserID,
	}

	err := ps.incomeRepo.Update(incomeToUpdate)
	if err != nil {
		return income, err
	}

	result, err := ps.incomeRepo.FindByID(income.ID)
	if err != nil {
		return nil, err
	}

	foundIncome := &types.IncomeDTO{
		ID:          result.ID,
		Description: result.Description,
		Amount:      result.Amount,
		Category:    result.Category,
		CreatedAt:   result.CreatedAt,
		UserID:      result.UserID,
	}

	return foundIncome, nil
}

func (ps *incomeServices) Delete(id uuid.UUID) error {
	err := ps.incomeRepo.Delete(id)
	return err
}
