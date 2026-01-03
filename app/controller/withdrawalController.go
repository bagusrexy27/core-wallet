package controller

import (
	"github.com/core-wallet/app/dto"
	"github.com/core-wallet/app/handler"
	"github.com/core-wallet/app/utils"
	"github.com/gofiber/fiber/v2"
)

type WithdrawalController struct {
	withdrawalHandler *handler.WithdrawalHandler
}

func NewWithdrawalController(
	withdrawalHandler *handler.WithdrawalHandler,
) *WithdrawalController {
	return &WithdrawalController{
		withdrawalHandler: withdrawalHandler,
	}
}

func (c *WithdrawalController) RequestWithdrawal(ctx *fiber.Ctx) error {
	var request dto.WithdrawalRequest

	if err := ctx.BodyParser(&request); err != nil {
		utils.ErrorLog("invalid body request", err, false)
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	trxId, err := c.withdrawalHandler.RequestWithdrawal(ctx.Context(), request)
	if err != nil {
		if err.Error() == "insufficient balance" {
			utils.ErrorLog("insufficient balance", err, false)
			return utils.ResponseError(ctx, fiber.StatusBadRequest, err.Error())
		}
		utils.ErrorLog("something went wrong", err, false)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return utils.ResponseSuccess(ctx, trxId, "request withdrawal berhasil")
}

func (c *WithdrawalController) ConfirmWithdrawal(ctx *fiber.Ctx) error {
	var request dto.ConfirmWithdrawalRequest

	if err := ctx.BodyParser(&request); err != nil {
		utils.ErrorLog("invalid body request", err, false)
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	if err := c.withdrawalHandler.ConfirmWithdrawal(ctx.Context(), request); err != nil {
		utils.ErrorLog("something went wrong", err, false)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "something went wrong")
	}
	return utils.ResponseSuccess(ctx, nil, "confirm withdrawal berhasil")
}

func (c *WithdrawalController) CheckStatusWithdrawal(ctx *fiber.Ctx) error {
	var request dto.CheckStatusTopUpRequest

	if err := ctx.BodyParser(&request); err != nil {
		utils.ErrorLog("invalid body request", err, false)
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	trx, err := c.withdrawalHandler.CheckStatusTransaction(ctx.Context(), request)
	if err != nil {
		utils.ErrorLog("something went wrong", err, false)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "something went wrong")
	}
	return utils.ResponseSuccess(ctx, trx, "check status withdrawal berhasil")
}
