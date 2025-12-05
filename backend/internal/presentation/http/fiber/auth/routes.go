package auth

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/common/registry"
)

type authRoutes struct {
	ctrl   Controller
	router http.HTTPRouter
}

func NewRoutes(router http.HTTPRouter, ctrl Controller) registry.Routes {
	return &authRoutes{
		ctrl:   ctrl,
		router: router,
	}
}

func (ar *authRoutes) Register() {
	authRoutes := ar.router.Group("/auth")
	authRoutes.Post("/users", ar.ctrl.Register)
	authRoutes.Post("/sessions", ar.ctrl.Login)
}
