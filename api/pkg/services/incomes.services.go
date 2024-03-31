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
		ID:        income.ID,
		Amount:    income.Amount,
		Category:  income.Category,
		CreatedAt: income.CreatedAt,
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
		ID:        result.ID,
		Amount:    result.Amount,
		Category:  result.Category,
		CreatedAt: result.CreatedAt,
	}
	return foundIncome, nil
}

func (ps *incomeServices) FindAll() ([]*types.IncomeDTO, error) {
	result, err := ps.incomeRepo.FindAll()
	if err != nil {
		return nil, err
	}

	foundIncomesList := []*types.IncomeDTO{}

	for _, income := range result {
		parse := &types.IncomeDTO{
			ID:        income.ID,
			Amount:    income.Amount,
			Category:  income.Category,
			CreatedAt: income.CreatedAt,
		}

		foundIncomesList = append(foundIncomesList, parse)
	}

	return foundIncomesList, nil
}

func (ps *incomeServices) Update(income *types.IncomeDTO, updates interface{}) (*types.IncomeDTO, error) {
	err := ps.incomeRepo.Update(income.ID, updates)
	if err != nil {
		return income, err
	}

	result, err := ps.incomeRepo.FindByID(income.ID)
	if err != nil {
		return nil, err
	}

	foundIncome := &types.IncomeDTO{
		ID:        result.ID,
		Amount:    result.Amount,
		Category:  result.Category,
		CreatedAt: result.CreatedAt,
	}

	return foundIncome, nil
}
