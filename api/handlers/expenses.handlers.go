package handlers

import (
	"errors"
	"fmt"
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/types"
	"gofinance/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateExpense(expenseService interfaces.ExpenseServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		b := new(types.ExpenseDTO)

		if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error parsing body: %v", err))
		}

		authHeader := ctx.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		tokenInfo, err := utils.ParseToken(tokenString)
		if err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("failed to parse token: %v", err))
		}

		b.UserID = tokenInfo.ID
		b.ID = uuid.New()
		b.CreatedAt = time.Now()

		result, err := expenseService.Create(b)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.JSON(result)
	}
}

func UpdateExpenses(expensesService interfaces.ExpenseServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(errors.New("invalid expense ID"))
		}

		b := new(types.ExpenseDTO)
		if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error parsing body: %v", err))
		}

		updatedExpense, err := expensesService.Update(b, id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.JSON(updatedExpense)
	}
}

func DeleteExpenses(expensesService interfaces.ExpenseServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(errors.New("invalid expense ID"))
		}

		err = expensesService.Delete(id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.SendStatus(http.StatusNoContent)
	}
}
