package routes

import (
	"gofinance/api/handlers"
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func ExpensesRoutes(app fiber.Router, expenseService interfaces.ExpenseServices) {
	r := app.Group("/expenses", handlers.TokenVerificationMiddleware(services.InitUsersServices()))

	// r.Get("/", handlers.ExpensesList)
	r.Post("/", handlers.CreateExpense(expenseService))
	// r.Get("/:id", handlers.ExpensesRead)
	// r.Put("/:id", handlers.ExpensesUpdate)
	// r.Delete("/:id", handlers.ExpensesDelete)
}
