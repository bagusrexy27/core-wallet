package controller

import (
	"net/http"

	"github.com/core-wallet/app/handler"
	"github.com/core-wallet/app/utils"
	"github.com/gofiber/fiber/v2"
)

type CheckHealthController struct {
	handler handler.CheckHealthHandler
}

func NewCheckHealthController(h handler.CheckHealthHandler) *CheckHealthController {
	return &CheckHealthController{
		handler: h,
	}
}

func (c *CheckHealthController) CheckHealth(ctx *fiber.Ctx) error {
	if err := c.handler.CheckHealth(); err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, "API is not healthy")
		return err
	}
	utils.ResponseSuccess(ctx, nil, "API is healthy")
	return nil
}
