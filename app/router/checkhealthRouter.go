package router

import (
	"github.com/core-wallet/app/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterCheckHealthRoutes(router fiber.Router, h *controller.CheckHealthController) {
	routeCheckhealth := router.Group("/health")
	routeCheckhealth.Get("/", h.CheckHealth)
}
