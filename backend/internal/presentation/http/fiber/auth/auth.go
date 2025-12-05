package auth

import (
	authDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	httpAbstraction "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"
	httperrors "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http/errors"
	"net/http"
)

type controller struct {
	service authDomain.Service
	lgr     logger.Logger
}

type Controller interface {
	Register(ctx httpAbstraction.HTTPContext) error
	Login(ctx httpAbstraction.HTTPContext) error
}

func NewController(service authDomain.Service, lgr logger.Logger) Controller {
	return &controller{
		service: service,
		lgr:     lgr,
	}
}

func (c *controller) Register(ctx httpAbstraction.HTTPContext) error {
	var cmd authDomain.RegisterCommand
	if err := ctx.BindJSON(&cmd); err != nil {
		return httperrors.ErrBadRequest
	}

	result, err := c.service.Register(ctx.Context(), cmd)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(result)
}

func (c *controller) Login(ctx httpAbstraction.HTTPContext) error {
	var cmd authDomain.LoginCommand
	if err := ctx.BindJSON(&cmd); err != nil {
		return httperrors.ErrBadRequest
	}

	result, err := c.service.Login(ctx.Context(), cmd)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(result)
}
