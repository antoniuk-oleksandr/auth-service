package errorhandler

import (
	authDomain "github.com/antoniuk-oleksandr/auth-service/frontend/internal/domain/auth"

	"github.com/gofiber/fiber/v2"
)

var errorMethods = map[error]int{
	authDomain.ErrInvalidCredentials:   fiber.StatusUnauthorized,
	authDomain.ErrFailedToLoginUser:    fiber.StatusInternalServerError,
	authDomain.ErrFailedToRegisterUser: fiber.StatusInternalServerError,
}
