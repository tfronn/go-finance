package routes

import (
	"gofinance/api/handlers"
	"gofinance/api/pkg/interfaces"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app fiber.Router, userService interfaces.UserServices) {
	r := app.Group("/users")

	r.Post("/", handlers.CreateUser(userService))
	r.Post("/login", handlers.LoginByEmailAndPassword(userService))
	r.Get("/google-login", handlers.GoogleLogin(userService))
	r.Get("/", handlers.TokenVerificationMiddleware(userService), handlers.GetUserInfoByToken(userService))
}
