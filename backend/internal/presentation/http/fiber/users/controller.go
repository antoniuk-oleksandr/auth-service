package users

import (
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"
	httperrors "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http/errors"

	"github.com/gofiber/fiber/v2"
)

type controller struct {
	service usersDomain.Service
	lgr     logger.Logger
}

type Controller interface {
	GetUserByID(ctx http.HTTPContext) error
}

func NewController(service usersDomain.Service, lgr logger.Logger) Controller {
	return &controller{
		service: service,
		lgr:     lgr,
	}
}

func (c *controller) GetUserByID(ctx http.HTTPContext) error {
	id := ctx.Param("id")
	if id == "" {
		return httperrors.ErrBadRequest
	}

	user, err := c.service.GetUserByID(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
