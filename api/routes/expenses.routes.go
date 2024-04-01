package routes

import (
	"gofinance/api/handlers"
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func ExpensesRoutes(app fiber.Router, expenseService interfaces.ExpenseServices) {
	r := app.Group("/expenses", handlers.TokenVerificationMiddleware(services.InitUsersServices()))

	r.Post("/", handlers.CreateExpense(expenseService))
	r.Put("/:id", handlers.UpdateExpenses(expenseService))
	r.Delete("/:id", handlers.DeleteExpenses(expenseService))
}
