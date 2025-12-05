package fxmodules

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/config"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"github.com/antoniuk-oleksandr/auth-service/backend/pkg/hasher"
	uuidgeneator "github.com/antoniuk-oleksandr/auth-service/backend/pkg/uuid"
	commoncaarlos0_env "github.com/antoniuk-oleksandr/auth-service/common/env_parser/caarlos0_env"
	commonvalidatorv10 "github.com/antoniuk-oleksandr/auth-service/common/validator/validatorv10"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var CommonModule = fx.Module("common",
	fx.Invoke(loadEnv),
	fx.Provide(
		commoncaarlos0_env.New,
		commonvalidatorv10.New,
		config.LoadAppConfig,
		provideLogger,
		uuidgeneator.New,
		provideBcryptHasher,
	),
)

func loadEnv() {
	_ = godotenv.Load()
}

func provideLogger(cfg *config.AppConfig) (logger.Logger, error) {
	return logger.New(cfg.Logger.Type)
}

func provideBcryptHasher(cfg *config.AppConfig) hasher.Hasher {
	return hasher.NewBcryptHasher(cfg.Hasher.Cost)
}
