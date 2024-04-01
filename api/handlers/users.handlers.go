package handlers

import (
	"net/http"
	"strings"

	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/types"
	"gofinance/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(userService interfaces.UserServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userDTO := &types.UserDTO{}
		if err := c.BodyParser(userDTO); err != nil {
			return err
		}

		user, err := userService.Create(userDTO)
		if err != nil {
			return err
		}

		return c.JSON(user)
	}
}

func GetUserInfoByToken(userService interfaces.UserServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": http.StatusUnauthorized,
				"error":      "Missing authorization token",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userDTO, err := userService.FindByToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": http.StatusUnauthorized,
				"error":      err.Error(),
			})
		}

		return c.JSON(userDTO)
	}
}

func GoogleLogin(userService interfaces.UserServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		googleID := c.Query("googleID")
		if googleID == "" {
			return fiber.NewError(http.StatusBadRequest, "Google ID is required")
		}

		user, err := userService.FindByGoogleID(googleID)
		if err != nil {
			return err
		}

		userDTO := &types.UserDTO{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: "",
			GoogleID: user.GoogleID,
		}

		token, err := utils.GenerateToken(userDTO)
		if err != nil {
			return err
		}

		return c.JSON(map[string]interface{}{
			"user":   userDTO,
			"token":  token,
			"status": "success",
		})
	}
}

func TokenVerificationMiddleware(userService interfaces.UserServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": http.StatusUnauthorized,
				"error":      "Missing authorization token",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userDTO, err := userService.FindByToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": http.StatusUnauthorized,
				"error":      err.Error(),
			})
		}

		c.Context().SetUserValue("user", userDTO)

		return c.Next()
	}
}
