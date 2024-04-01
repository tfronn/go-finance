package services

import (
	"fmt"
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/models"
	"gofinance/api/pkg/types"
	"gofinance/utils"
	"time"

	"github.com/google/uuid"
)

type userServices struct {
	userRepo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserServices {
	return &userServices{userRepo: repo}
}

func (us *userServices) Create(user *types.UserDTO) (*types.UserDTO, error) {
	newUser := &models.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		GoogleID: user.GoogleID,
	}

	_, err := us.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userServices) FindByID(id uuid.UUID) (*types.UserDTO, error) {
	result, err := us.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	expensesDTO := make([]*types.ExpenseDTO, len(result.Expenses))
	for i, expense := range result.Expenses {
		expensesDTO[i] = &types.ExpenseDTO{
			ID:        expense.ID,
			Amount:    expense.Amount,
			CreatedAt: expense.CreatedAt,
			Category:  expense.Category,
			UserID:    expense.UserID,
		}
	}

	incomesDTO := make([]*types.IncomeDTO, len(result.Incomes))
	for i, income := range result.Incomes {
		incomesDTO[i] = &types.IncomeDTO{
			ID:        income.ID,
			Amount:    income.Amount,
			CreatedAt: income.CreatedAt,
			Category:  income.Category,
			UserID:    income.UserID,
		}
	}

	userDTO := &types.UserDTO{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Password:  result.Password,
		GoogleID:  result.GoogleID,
		CreatedAt: result.CreatedAt,
		Incomes:   incomesDTO,
		Expenses:  expensesDTO,
	}

	return userDTO, nil
}

func (us *userServices) FindAll() ([]*types.UserDTO, error) {
	users, err := us.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	foundUsers := []*types.UserDTO{}

	for _, user := range users {
		userDTO := &types.UserDTO{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			GoogleID: user.GoogleID,
		}
		foundUsers = append(foundUsers, userDTO)
	}

	return foundUsers, nil
}

func (us *userServices) FindByToken(tokenString string) (*types.UserDTO, error) {
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	result, err := us.FindByID(claims.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	// Check if the token email matches the user email
	if result.Email != claims.Email {
		return nil, fmt.Errorf("token email doesn't match the user email")
	}

	// Check if the token has not expired
	if claims.ExpiredAt.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}

	expensesDTO := make([]*types.ExpenseDTO, len(result.Expenses))
	for i, expense := range result.Expenses {
		expensesDTO[i] = &types.ExpenseDTO{
			ID:        expense.ID,
			Amount:    expense.Amount,
			CreatedAt: expense.CreatedAt,
			Category:  expense.Category,
			UserID:    expense.UserID,
		}
	}

	incomesDTO := make([]*types.IncomeDTO, len(result.Incomes))
	for i, income := range result.Incomes {
		incomesDTO[i] = &types.IncomeDTO{
			ID:        income.ID,
			Amount:    income.Amount,
			CreatedAt: income.CreatedAt,
			Category:  income.Category,
			UserID:    income.UserID,
		}
	}

	userDTO := &types.UserDTO{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Password:  result.Password,
		GoogleID:  result.GoogleID,
		CreatedAt: result.CreatedAt,
		Incomes:   incomesDTO,
		Expenses:  expensesDTO,
	}

	return userDTO, nil

}

func (us *userServices) FindByEmail(email string) (*types.UserDTO, error) {
	result, err := us.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	expensesDTO := make([]*types.ExpenseDTO, len(result.Expenses))
	for i, expense := range result.Expenses {
		expensesDTO[i] = &types.ExpenseDTO{
			ID:        expense.ID,
			Amount:    expense.Amount,
			CreatedAt: expense.CreatedAt,
			Category:  expense.Category,
			UserID:    expense.UserID,
		}
	}

	incomesDTO := make([]*types.IncomeDTO, len(result.Incomes))
	for i, income := range result.Incomes {
		incomesDTO[i] = &types.IncomeDTO{
			ID:        income.ID,
			Amount:    income.Amount,
			CreatedAt: income.CreatedAt,
			Category:  income.Category,
			UserID:    income.UserID,
		}
	}

	userDTO := &types.UserDTO{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Password:  result.Password,
		GoogleID:  result.GoogleID,
		CreatedAt: result.CreatedAt,
		Incomes:   incomesDTO,
		Expenses:  expensesDTO,
	}

	return userDTO, nil
}

func (us *userServices) FindByGoogleID(googleID string) (*types.UserDTO, error) {
	result, err := us.userRepo.FindByGoogleID(googleID)
	if err != nil {
		return nil, err
	}

	expensesDTO := make([]*types.ExpenseDTO, len(result.Expenses))
	for i, expense := range result.Expenses {
		expensesDTO[i] = &types.ExpenseDTO{
			ID:        expense.ID,
			Amount:    expense.Amount,
			CreatedAt: expense.CreatedAt,
			Category:  expense.Category,
			UserID:    expense.UserID,
		}
	}

	incomesDTO := make([]*types.IncomeDTO, len(result.Incomes))
	for i, income := range result.Incomes {
		incomesDTO[i] = &types.IncomeDTO{
			ID:        income.ID,
			Amount:    income.Amount,
			CreatedAt: income.CreatedAt,
			Category:  income.Category,
			UserID:    income.UserID,
		}
	}

	userDTO := &types.UserDTO{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Password:  result.Password,
		GoogleID:  result.GoogleID,
		CreatedAt: result.CreatedAt,
		Incomes:   incomesDTO,
		Expenses:  expensesDTO,
	}

	return userDTO, nil
}

func (us *userServices) Update(user *types.UserDTO, updates interface{}) (*types.UserDTO, error) {
	err := us.userRepo.Update(user.ID, updates)
	if err != nil {
		return user, err
	}

	result, err := us.userRepo.FindByID(user.ID)
	if err != nil {
		return nil, err
	}

	expensesDTO := make([]*types.ExpenseDTO, len(result.Expenses))
	for i, expense := range result.Expenses {
		expensesDTO[i] = &types.ExpenseDTO{
			ID:        expense.ID,
			Amount:    expense.Amount,
			CreatedAt: expense.CreatedAt,
			Category:  expense.Category,
			UserID:    expense.UserID,
		}
	}

	incomesDTO := make([]*types.IncomeDTO, len(result.Incomes))
	for i, income := range result.Incomes {
		incomesDTO[i] = &types.IncomeDTO{
			ID:        income.ID,
			Amount:    income.Amount,
			CreatedAt: income.CreatedAt,
			Category:  income.Category,
			UserID:    income.UserID,
		}
	}

	userDTO := &types.UserDTO{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Password:  result.Password,
		GoogleID:  result.GoogleID,
		CreatedAt: result.CreatedAt,
		Incomes:   incomesDTO,
		Expenses:  expensesDTO,
	}

	return userDTO, nil

}
