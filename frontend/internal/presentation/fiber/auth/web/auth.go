package auth

import (
	common "github.com/antoniuk-oleksandr/auth-service/frontend/internal/presentation/fiber/common"
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/view/page/home"
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/view/page/login"
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/view/page/me"
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/view/page/register"

	"github.com/gofiber/fiber/v2"
)

type WebController interface {
	ShowRegisterPage(ctx *fiber.Ctx) error
	ShowLoginPage(ctx *fiber.Ctx) error
	ShowMePage(ctx *fiber.Ctx) error
	ShowHomePage(ctx *fiber.Ctx) error
}

type controller struct {
}

func NewController() WebController {
	return &controller{}
}

func (c *controller) ShowHomePage(ctx *fiber.Ctx) error {
	return common.Page(ctx, home.HomePage())
}

func (c *controller) ShowLoginPage(ctx *fiber.Ctx) error {
	return common.Page(ctx, login.LoginPage())
}

func (c *controller) ShowMePage(ctx *fiber.Ctx) error {
	return common.Page(ctx, me.MePage())
}

func (c *controller) ShowRegisterPage(ctx *fiber.Ctx) error {
	return common.Page(ctx, register.RegisterPage())
}
