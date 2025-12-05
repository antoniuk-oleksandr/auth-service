package fxmodules

import (
	authApp "github.com/antoniuk-oleksandr/auth-service/frontend/internal/application/auth"

	"go.uber.org/fx"
)

var AuthAppModule = fx.Module("auth",
	fx.Provide(authApp.NewService),
)
