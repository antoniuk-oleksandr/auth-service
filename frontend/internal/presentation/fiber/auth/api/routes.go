package auth

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, ctrl APIController) {
	g := r.Group("auth")

	g.Post("/sessions", ctrl.LoginUser)
	g.Post("/users", ctrl.RegisterUser)
}
