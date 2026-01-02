package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *fiber.Ctx, data interface{}, message string) error {
	resp := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func ResponseError(c *fiber.Ctx, httpCode int, message string) error {
	return c.Status(httpCode).JSON(Response{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}
