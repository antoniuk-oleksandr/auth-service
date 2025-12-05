package auth

import (
	authDomain "github.com/antoniuk-oleksandr/auth-service/frontend/internal/domain/auth"
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/presentation/fiber/common"

	"github.com/gofiber/fiber/v2"
)

type APIController interface {
	LoginUser(ctx *fiber.Ctx) error
	RegisterUser(ctx *fiber.Ctx) error
}

type controller struct {
	srv       authDomain.AuthService
	responder AuthResponder
}

func NewController(srv authDomain.AuthService, r AuthResponder) APIController {
	return &controller{
		srv:       srv,
		responder: r,
	}
}

func (ctrl *controller) LoginUser(ctx *fiber.Ctx) error {
	var credentials authDomain.Credentials

	err := ctx.BodyParser(&credentials)
	if err != nil {
		common.TriggerError(ctx, "login:error", authDomain.ErrInvalidCredentials.Error())
		return authDomain.ErrInvalidCredentials
	}

	jwt, err := ctrl.srv.LoginUser(&credentials)
	if err != nil {
		common.TriggerError(ctx, "login:error", err.Error())
		return err
	}

	return ctrl.responder.Send(ctx, jwt, "login:success")
}

func (ctrl *controller) RegisterUser(ctx *fiber.Ctx) error {
	var credentials authDomain.Credentials
	err := ctx.BodyParser(&credentials)
	if err != nil {
		common.TriggerError(ctx, "login:error", authDomain.ErrInvalidCredentials.Error())
		return authDomain.ErrInvalidCredentials
	}

	jwt, err := ctrl.srv.RegisterUser(&credentials)
	if err != nil {
		common.TriggerError(ctx, "login:error", err.Error())
		return err
	}

	return ctrl.responder.Send(ctx, jwt, "register:success")
}
