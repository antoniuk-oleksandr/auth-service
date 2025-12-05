package errorhandler

import (
	"github.com/gofiber/fiber/v2"
)

func New(ctx *fiber.Ctx, err error) error {
	if status, ok := errorMethods[err]; ok {
		return ctx.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": fiber.ErrInternalServerError.Message,
	})
}
