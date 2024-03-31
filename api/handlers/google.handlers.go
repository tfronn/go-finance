package handlers

import (
	"context"
	"encoding/json"
	"io"

	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/services"
	"gofinance/api/pkg/types"
	"gofinance/utils"

	"github.com/gofiber/fiber/v2"
)

func GoogleOAuth2Login(userService interfaces.UserServices) fiber.Handler {
	conf := services.NewGoogleOAuth2Config()

	return func(c *fiber.Ctx) error {
		url := conf.AuthCodeURL("state")
		return c.Redirect(url)
	}
}

func GoogleOAuth2Callback(userService interfaces.UserServices) fiber.Handler {
	conf := services.NewGoogleOAuth2Config()

	return func(c *fiber.Ctx) error {
		code := c.Query("code")
		if code == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Missing code parameter")
		}

		token, err := conf.Exchange(context.Background(), code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange code for token")
		}

		client := conf.Client(context.Background(), token)
		response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user's email address")
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to read response body")
		}
		defer response.Body.Close()

		// Do something with the body, for example, parse it
		var userInfo map[string]interface{}
		err = json.Unmarshal(body, &userInfo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to parse user information")
		}

		// Find or create a user based on the email address
		email := userInfo["email"].(string)
		user, err := userService.FindByEmail(email)
		if err != nil {
			// Create a new user if it doesn't exist
			newUser := &types.UserDTO{
				Email: email,
				// Set other user properties
			}
			user, err = userService.Create(newUser)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to create user")
			}
		}

		generatedToken, err := utils.GenerateToken(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate token")
		}

		return c.JSON(fiber.Map{
			"user":   user,
			"token":  generatedToken,
			"status": "success",
		})
	}
}
