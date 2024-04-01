package main

import (
	"fmt"
	"gofinance/api/pkg/models"
	"gofinance/api/pkg/services"
	"gofinance/api/routes"
	"gofinance/config"
	"gofinance/config/database"
	"gofinance/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.Migrate(
		&models.User{},
		&models.Income{},
		&models.Expense{},
	)

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	app.Use(cors.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("Everything sounds good...")
	})
	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Z--AppName", "GoFinance")
		c.Set("Z--APIVersion", "v1")
		return c.Next()
	})

	routes.IncomesRoutes(v1, services.InitIncomeServices())
	routes.ExpensesRoutes(v1, services.InitExpenseServices())
	routes.UsersRoutes(v1, services.InitUsersServices())
	routes.GoogleOAuth2Routes(v1, services.InitUsersServices())

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
