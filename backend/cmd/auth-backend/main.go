package main

import (
	fxmodules "github.com/antoniuk-oleksandr/auth-service/backend/internal/fx"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fxmodules.CommonModule,
		fxmodules.DBModule,
		fxmodules.RepoModule,
		fxmodules.UsersAppModule,
		fxmodules.AuthAppModule,
		// fxmodules.HTTPModule,
		fxmodules.CTPModule,
	)

	app.Run()
}
