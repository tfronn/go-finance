package routes

import (
	"gofinance/api/handlers"
	"gofinance/api/pkg/interfaces"

	"github.com/gofiber/fiber/v2"
)

func GoogleOAuth2Routes(app fiber.Router, userService interfaces.UserServices) {
	r := app.Group("/auth/google")

	r.Get("/login", handlers.GoogleOAuth2Login(userService))
	r.Get("/callback", handlers.GoogleOAuth2Callback(userService))
}
