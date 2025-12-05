package fxmodules

import (
	usersApp "github.com/antoniuk-oleksandr/auth-service/backend/internal/application/users"

	"go.uber.org/fx"
)

var UsersAppModule = fx.Module("users",
	fx.Provide(usersApp.NewService),
)
