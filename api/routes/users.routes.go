package routes

import (
	"gofinance/api/handlers"
	"gofinance/api/pkg/interfaces"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app fiber.Router, userService interfaces.UserServices) {
	r := app.Group("/users")

	r.Post("/", handlers.CreateUser(userService))
	// r.Get("/:id", handlers.GetUserByID(userService))
	// r.Get("/", handlers.GetUsers(userService))

	// Google Login route
	r.Get("/google-login", handlers.GoogleLogin(userService))
}
