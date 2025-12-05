package fxmodules

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/config"
	routesRegistry "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/common/registry"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"
	authHttp "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http/fiber/auth"
	usersHttp "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http/fiber/users"
	router "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http/router"
	"fmt"

	"go.uber.org/fx"
)

var HTTPModule = fx.Module("http",
	fx.Provide(provideRouter),
	fx.Provide(routesRegistry.New),

	fx.Provide(usersHttp.NewController),
	fx.Provide(authHttp.NewController),

	fx.Invoke(provideUsersRoutes),
	fx.Invoke(provideAuthRoutes),
	fx.Invoke(registerHTTPRoutes),

	fx.Invoke(startServer),
)

func provideRouter(cfg *config.AppConfig) (http.HTTPRouter, error) {
	switch cfg.Server.HTTPFramework {
	case "fiber":
		return router.NewHTTPRouter(cfg.Server.HTTPFramework, config.NewFiberConfig())
	default:
		return nil, fmt.Errorf("unsupported HTTP framework: %s", cfg.Server.HTTPFramework)
	}
}

func startServer(r http.HTTPRouter, cfg *config.AppConfig) error {
	return r.Start(":" + cfg.Server.Port)
}

func provideUsersRoutes(
	registry routesRegistry.Registry,
	ctrl usersHttp.Controller,
	r http.HTTPRouter,
) {
	registry.Add(usersHttp.NewRoutes(r, ctrl))
}

func provideAuthRoutes(
	registry routesRegistry.Registry,
	ctrl authHttp.Controller,
	r http.HTTPRouter,
) {
	registry.Add(authHttp.NewRoutes(r, ctrl))
}

func registerHTTPRoutes(
	registry routesRegistry.Registry,
) {
	registry.Register()
}
