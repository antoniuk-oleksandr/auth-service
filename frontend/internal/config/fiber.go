package config

import (
	errorhandler "github.com/antoniuk-oleksandr/auth-service/frontend/internal/presentation/fiber/error_handler/fiber"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: errorhandler.New,
	}
}
