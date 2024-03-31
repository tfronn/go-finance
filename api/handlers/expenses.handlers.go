package handlers

import (
	"fmt"
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/types"
	"gofinance/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateExpense(expenseService interfaces.ExpenseServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		b := new(types.ExpenseDTO)

		if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error parsing body: %v", err))
		}

		result, err := expenseService.Create(b)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.JSON(result)
	}
}
