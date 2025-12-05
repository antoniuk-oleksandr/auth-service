package auth

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, ctrl WebController) {
	r.Get("/login", ctrl.ShowLoginPage)
	r.Get("/register", ctrl.ShowRegisterPage)
	r.Get("/me", ctrl.ShowMePage)
	r.Get("/", ctrl.ShowHomePage)
}