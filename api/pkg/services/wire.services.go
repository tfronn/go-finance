package services

import (
	"gofinance/api/pkg/interfaces"
	repository "gofinance/api/pkg/repositories"
	"gofinance/config/database"
)

func InitIncomeServices() interfaces.IncomeServices {
	incomeRepo := repository.NewIncomeRepository(database.DB)
	incomeServices := NewIncomeService(incomeRepo)

	return incomeServices
}

func InitExpenseServices() interfaces.ExpenseServices {
	expenseRepo := repository.NewExpenseRepository(database.DB)
	expenseServices := NewExpenseService(expenseRepo)

	return expenseServices
}

func InitUsersServices() interfaces.UserServices {
	usersRepo := repository.NewUserRepository(database.DB)
	usersServices := NewUserService(usersRepo)

	return usersServices
}
