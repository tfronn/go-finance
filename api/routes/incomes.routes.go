package routes

import (
	"gofinance/api/handlers"
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func IncomesRoutes(app fiber.Router, incomeService interfaces.IncomeServices) {
	r := app.Group("/incomes", handlers.TokenVerificationMiddleware(services.InitUsersServices()))

	r.Post("/", handlers.CreateIncome(incomeService))
	r.Put("/:id", handlers.UpdateIncome(incomeService))
	r.Delete("/:id", handlers.IncomesDelete(incomeService))
}
