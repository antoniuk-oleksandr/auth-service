package fxmodules

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/config"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	ctplogger "github.com/antoniuk-oleksandr/auth-service/backend/internal/logger/ctp_logger"
	routesRegistry "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/common/registry"
	authCTP "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/ctp/auth"
	usersCTP "github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/ctp/users"
	"github.com/antoniuk-oleksandr/auth-service/ctp/server"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"

	"go.uber.org/fx"
)

var CTPModule = fx.Module("ctp",
	fx.Provide(provideCTPServer),
	fx.Provide(routesRegistry.New),

	fx.Provide(authCTP.NewController),
	fx.Provide(usersCTP.NewController),

	fx.Invoke(provideAuthCTPRoutes),
	fx.Invoke(provideUsersCTPRoutes),
	fx.Invoke(registerCTPRoutes),

	fx.Invoke(startCTPServer),
)

func provideUsersCTPRoutes(
	registry routesRegistry.Registry,
	ctrl usersCTP.Controller,
	conn types.Server,
) {
	registry.Add(usersCTP.NewRoutes(conn, ctrl))
}

func provideAuthCTPRoutes(
	registry routesRegistry.Registry,
	ctrl authCTP.Controller,
	conn types.Server,
) {
	registry.Add(authCTP.NewRoutes(conn, ctrl))
}

func registerCTPRoutes(
	registry routesRegistry.Registry,
) {
	registry.Register()
}

func startCTPServer(srv types.Server, cfg *config.AppConfig) error {
	return srv.Start(":" + cfg.Server.Port)
}

func provideCTPServer(logger logger.Logger) types.Server {
	lgr := ctplogger.NewCTPLogger(logger)
	opt := server.WithLogger(lgr)
	return server.NewServer(opt)
}
