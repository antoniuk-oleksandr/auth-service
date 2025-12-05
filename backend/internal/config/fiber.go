package config

import (
	errorhandler "github.com/antoniuk-oleksandr/auth-service/backend/internal/error_handler/http/fiber"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: errorhandler.New,
	}
}
