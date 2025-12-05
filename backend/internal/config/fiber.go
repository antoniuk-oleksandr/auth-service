package config

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	errorhandler "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http/fiber/error_handler/http/fiber"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig(lgr logger.Logger) fiber.Config {
	return fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return errorhandler.New(ctx, err, lgr)
		},
	}
}
