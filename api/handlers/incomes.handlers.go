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

func CreateIncome(incomeService interfaces.IncomeServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		b := new(types.IncomeDTO)
		authHeader := ctx.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		tokenInfo, err := utils.ParseToken(tokenString)
		if err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("failed to parse token: %v", err))
		}

		if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error parsing body: %v", err))
		}

		b.UserID = tokenInfo.ID
		b.ID = uuid.New()
		b.CreatedAt = time.Now()

		result, err := incomeService.Create(b)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.JSON(result)
	}
}

func UpdateIncome(incomesService interfaces.IncomeServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(errors.New("invalid income ID"))
		}

		b := &types.IncomeDTO{}
		if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error parsing body: %v", err))
		}

		income, err := incomesService.FindByID(id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}

		b.ID = income.ID
		b.CreatedAt = income.CreatedAt
		b.UserID = income.UserID

		updatedIncome, err := incomesService.Update(b, id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err.Error())
		}
		return ctx.JSON(updatedIncome)
	}
}

func DeleteIncome(incomesService interfaces.IncomeServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(errors.New("invalid income ID"))
		}

		err = incomesService.Delete(id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.SendStatus(http.StatusNoContent)
	}
}
