package fxmodules

import (
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/config"

	authAPI "github.com/antoniuk-oleksandr/auth-service/frontend/internal/presentation/fiber/auth/api"
	authWeb "github.com/antoniuk-oleksandr/auth-service/frontend/internal/presentation/fiber/auth/web"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type (
	Routers struct {
		fx.Out
		APIRouter fiber.Router `name:"api"`
		WebRouter fiber.Router `name:"web"`
	}

	APIRouteDeps struct {
		fx.In
		Router fiber.Router `name:"api"`
		Ctrl   authAPI.APIController
	}

	WebRouteDeps struct {
		fx.In
		Router fiber.Router `name:"web"`
		Ctrl   authWeb.WebController
	}
)

var HTTPModule = fx.Module("http",
	fx.Provide(provideFiberApp),
	fx.Provide(provideRouters),

	fx.Provide(authAPI.NewAuthResponder),
	fx.Provide(authAPI.NewController),
	fx.Provide(authWeb.NewController),

	fx.Invoke(provideAuthAPIRoutes),
	fx.Invoke(provideAuthWebRoutes),
	fx.Invoke(provideStaticFiles),

	fx.Invoke(startFiberApp),
)

func provideFiberApp() *fiber.App {
	return fiber.New(config.NewFiberConfig())
}

func provideRouters(app *fiber.App) Routers {
	return Routers{
		APIRouter: app.Group("/api/v1"),
		WebRouter: app.Group("/"),
	}
}

func provideAuthAPIRoutes(d APIRouteDeps) {
	authAPI.RegisterRoutes(d.Router, d.Ctrl)
}

func provideAuthWebRoutes(d WebRouteDeps) {
	authWeb.RegisterRoutes(d.Router, d.Ctrl)
}

func provideStaticFiles(app *fiber.App) {
	app.Static("/static", "./internal/presentation/fiber/static")
}

func startFiberApp(app *fiber.App, cfg *config.AppConfig) error {
	return app.Listen(":" + cfg.ServerConfig.Port)
}
