package router

import (
	"github.com/core-wallet/app/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterWalletRoutes(router fiber.Router, h *controller.WalletController) {
	routeWallet := router.Group("/wallet")
	routeWallet.Post("/create", h.CreateWallet)
	routeWallet.Post("/topup/request", h.RequestTopUp)
	routeWallet.Post("/topup/confirm", h.ConfirmTopUp)
	routeWallet.Post("/topup/reject", h.RejectTopUp)
	routeWallet.Post("/topup/status", h.CheckStatusTopUp)
	routeWallet.Get("/balance/:walletID", h.GetUserBalance)

}
