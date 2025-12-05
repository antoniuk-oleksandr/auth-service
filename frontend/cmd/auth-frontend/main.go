package main

import (
	fxmodules "github.com/antoniuk-oleksandr/auth-service/frontend/internal/fx"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fxmodules.CommonMovdule,
		fxmodules.CTPModule,
		fxmodules.AuthAppModule,
		fxmodules.HTTPModule,
	)

	app.Run()
}