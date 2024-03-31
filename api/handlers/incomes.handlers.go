package handlers

import (
	"errors"
	"fmt"
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/types"
	"gofinance/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateIncome(incomeService interfaces.IncomeServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		b := new(types.IncomeDTO)

		if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error parsing body: %v", err))
		}

		result, err := incomeService.Create(b)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.JSON(result)
	}
}

func GetAllIncomes(incomesService interfaces.IncomeServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		incomes, err := incomesService.FindAll()
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.JSON(incomes)
	}
}

func GetIncomeByID(incomesService interfaces.IncomeServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(errors.New("invalid income ID"))
		}

		income, err := incomesService.FindByID(id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.JSON(income)
	}
}

func UpdateIncome(incomesService interfaces.IncomeServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(errors.New("invalid income ID"))
		}

		b := new(types.IncomeDTO)
		if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error parsing body: %v", err))
		}

		updatedIncome, err := incomesService.Update(b, id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}
		return ctx.JSON(updatedIncome)
	}
}

// func IncomesDelete(incomesService interfaces.IncomeServices) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 			id, err := uuid.Parse(ctx.Params("id"))
// 			if err != nil {
// 					ctx.Status(http.StatusBadRequest)
// 					return ctx.JSON(errors.New("invalid income ID"))
// 			}

// 			err = incomesService.Delete(id)
// 			if err != nil {
// 					if err == interfaces.ErrIncomeNotFound {
// 							ctx.Status(http.StatusNotFound)
// 							return ctx.JSON(err)
// 					}
// 					ctx.Status(http.StatusInternalServerError)
// 					return ctx.JSON(err)
// 			}
// 			return ctx.SendStatus(http.StatusNoContent)
// 	}
// }
