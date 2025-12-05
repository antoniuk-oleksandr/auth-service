package users

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/common/registry"
)

type usersRoutes struct {
	ctrl   Controller
	router http.HTTPRouter
}

func NewRoutes(router http.HTTPRouter, ctrl Controller) registry.Routes {
	return &usersRoutes{
		ctrl: ctrl,
		router: router,
	}
}

func (ur *usersRoutes) Register() {
	usersRoutes := ur.router.Group("/users")
	usersRoutes.Get("/:id", ur.ctrl.GetUserByID)
}
