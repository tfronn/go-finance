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

		generatedToken, err := utils.GenerateToken(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate token")
		}

		user.Password = ""

		return c.JSON(fiber.Map{
			"user":   user,
			"token":  generatedToken,
			"status": "success",
		})
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

		userDTO.Password = ""

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

		userDTO.Password = ""

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

		userDTO.Password = ""

		return c.Next()
	}
}

func LoginByEmailAndPassword(userService interfaces.UserServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginDTO := &types.LoginDTO{}

		if err := c.BodyParser(loginDTO); err != nil {
			return err
		}

		user, err := userService.FindByEmail(loginDTO.Email)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": http.StatusUnauthorized,
				"error":      "Invalid email or password",
			})
		}

		if user.Password == loginDTO.Password {
			generatedToken, err := utils.GenerateToken(user)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate token")
			}

			user.Password = ""

			return c.JSON(fiber.Map{
				"user":   user,
				"token":  generatedToken,
				"status": "success",
			})
		}

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": http.StatusUnauthorized,
			"error":      "Invalid email or password",
		})
	}
}
