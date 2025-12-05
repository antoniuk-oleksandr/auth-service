package auth

import (
	authDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/ctp/common"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
	"errors"
)

type Controller interface {
	Register(ctx types.Ctx) error
	Login(ctx types.Ctx) error
}

type controller struct {
	srv authDomain.Service
	lgr logger.Logger
}

func NewController(srv authDomain.Service, lgr logger.Logger) Controller {
	return &controller{
		srv: srv,
		lgr: lgr,
	}
}

func (ctrl *controller) Login(ctx types.Ctx) error {
	var cmd authDomain.LoginCommand
	if err := common.ParseBody(ctx, &cmd); err != nil {
		return err
	}

	jwt, err := ctrl.srv.Login(ctx.Context(), cmd)
	if err != nil {
		if errors.Is(err, authDomain.ErrInvalidCredentials) {
			return ctx.Status(types.StatusUnauthorized).Send(types.Map{
				"error": err.Error(),
			})
		} else {
			return ctx.Status(types.StatusInternalError).Send(types.Map{
				"error": authDomain.ErrAuthFailed.Error(),
			})
		}
	}

	return ctx.Status(types.StatusCreated).Send(jwt)
}

func (ctrl *controller) Register(ctx types.Ctx) error {
	var cmd authDomain.RegisterCommand
	if err := common.ParseBody(ctx, &cmd); err != nil {
		return err
	}

	jwt, err := ctrl.srv.Register(ctx.Context(), cmd)
	if err != nil {
		return ctx.Status(types.StatusInternalError).Send(types.Map{
			"error": "could not register user",
		})
	}

	return ctx.Status(types.StatusCreated).Send(jwt)
}
