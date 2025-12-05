package auth

import (
	authDomain "github.com/antoniuk-oleksandr/auth-service/frontend/internal/domain/auth"
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/presentation/fiber/common"

	"github.com/gofiber/fiber/v2"
)

type AuthResponder interface {
	Send(ctx *fiber.Ctx, jwt *authDomain.JWT, trigger string) error
}

type authResponder struct{}

func NewAuthResponder() AuthResponder {
	return &authResponder{}
}

func (r *authResponder) Send(ctx *fiber.Ctx, jwt *authDomain.JWT, trigger string) error {
	common.HTMXTrigger(ctx, map[string]any{
		trigger: map[string]any{
			"access_token":             jwt.AccessToken,
			"access_access_expires_in": jwt.AccessExpiresIn,
		},
	})

	common.HTMXRedirect(ctx, "/")

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    jwt.RefreshToken,
		MaxAge:   jwt.RefreshExpiresIn,
		HTTPOnly: true,
		SameSite: "Strict",
	})

	return ctx.Status(fiber.StatusCreated).JSON(jwt)
}
