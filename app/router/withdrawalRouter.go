package router

import (
	"github.com/core-wallet/app/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterWithdrawalRoutes(router fiber.Router, h *controller.WithdrawalController) {
	routeWithdrawal := router.Group("/withdrawal")
	routeWithdrawal.Post("/request", h.RequestWithdrawal)
	routeWithdrawal.Post("/confirm", h.ConfirmWithdrawal)
	routeWithdrawal.Post("/status", h.CheckStatusWithdrawal)
}
