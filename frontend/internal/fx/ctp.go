package fxmodules

import (
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/config"
	"github.com/antoniuk-oleksandr/auth-service/ctp/client"

	"go.uber.org/fx"
)

var CTPModule = fx.Module("ctp",
	fx.Provide(provideCTPClient),
)

func provideCTPClient(cfg *config.AppConfig) client.Client {
	return client.NewClient(cfg.BackendConfig.Addr)
}
