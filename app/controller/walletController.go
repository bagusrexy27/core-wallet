package controller

import (
	"github.com/core-wallet/app/dto"
	"github.com/core-wallet/app/handler"
	"github.com/core-wallet/app/utils"
	"github.com/gofiber/fiber/v2"
)

type WalletController struct {
	WalletHandler *handler.WalletHandler
	TopUpHandler  *handler.TopUpHandler
}

func NewWalletController(
	handlerWallet *handler.WalletHandler,
	topUpHandler *handler.TopUpHandler,
) *WalletController {
	return &WalletController{
		WalletHandler: handlerWallet,
		TopUpHandler:  topUpHandler,
	}
}

func (c *WalletController) CreateWallet(ctx *fiber.Ctx) error {
	var req dto.CreateWalletRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	if req.UserID == "" {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	walletID, err := c.WalletHandler.CreateWallet(ctx.Context(), req.UserID)
	if err != nil {
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "failed to create wallet")
	}

	resp := dto.CreateWalletResponse{
		WalletID: walletID,
		UserID:   req.UserID,
		Balance:  0,
	}
	return utils.ResponseSuccess(ctx, resp, "Wallet created successfully")
}

func (c *WalletController) RequestTopUp(ctx *fiber.Ctx) error {
	var request dto.TopUpRequest

	if err := ctx.BodyParser(&request); err != nil {
		utils.ErrorLog("invalid body request", err, false)
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	trxId, err := c.TopUpHandler.TopUpRequest(ctx.Context(), request)
	if err != nil {
		utils.ErrorLog("something went wrong", err, false)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "something went wrong")
	}
	return utils.ResponseSuccess(ctx, fiber.Map{"trx_id": trxId}, "request top up berhasil")
}

func (c *WalletController) ConfirmTopUp(ctx *fiber.Ctx) error {
	var request dto.ConfirmTopUpRequest

	if err := ctx.BodyParser(&request); err != nil {
		utils.ErrorLog("invalid body request", err, false)
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	if err := c.TopUpHandler.ConfirmTopUp(ctx.Context(), request); err != nil {
		utils.ErrorLog("something went wrong", err, false)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "something went wrong")
	}
	return utils.ResponseSuccess(ctx, nil, "confirm top up berhasil")
}

func (c *WalletController) RejectTopUp(ctx *fiber.Ctx) error {
	var request dto.CheckStatusTopUpRequest

	if err := ctx.BodyParser(&request); err != nil {
		utils.ErrorLog("invalid body request", err, false)
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	if err := c.TopUpHandler.RejectTopUp(ctx.Context(), request); err != nil {
		utils.ErrorLog("something went wrong", err, false)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "something went wrong")
	}
	return utils.ResponseSuccess(ctx, nil, "reject top up berhasil")
}

func (c *WalletController) CheckStatusTopUp(ctx *fiber.Ctx) error {
	var request dto.CheckStatusTopUpRequest

	if err := ctx.BodyParser(&request); err != nil {
		utils.ErrorLog("invalid body request", err, false)
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "invalid request payload")
	}

	trx, err := c.TopUpHandler.CheckStatusTopUp(ctx.Context(), request)
	if err != nil {
		utils.ErrorLog("something went wrong", err, false)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "something went wrong")
	}
	return utils.ResponseSuccess(ctx, trx, "check status top up berhasil")
}

func (c WalletController) GetUserBalance(ctx *fiber.Ctx) error {
	walletID := ctx.Params("walletID")
	balance, err := c.WalletHandler.GetUserBalanceByWalletId(walletID)
	if err != nil {
		utils.ErrorLog("something went wrong", err, false)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "something went wrong")
	}
	return utils.ResponseSuccess(ctx, balance, "get balance berhasil")
}
