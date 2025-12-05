package fxmodules

import (
	commonconfig "auth-common/config"
	commonenvparser "auth-common/env_parser"
	commoncaarlos0_env "auth-common/env_parser/caarlos0_env"
	commonvalidator "auth-common/validator"
	commonvalidatorv10 "auth-common/validator/validatorv10"
	"github.com/antoniuk-oleksandr/auth-service/frontend/internal/config"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var CommonMovdule = fx.Module("common",
	fx.Invoke(loadEnv),
	fx.Provide(
		commoncaarlos0_env.New,
		commonvalidatorv10.New,
		proideAppConfig,
	),
)

func proideAppConfig(v commonvalidator.Validator, parser commonenvparser.EnvParser) *config.AppConfig {
	var cfg config.AppConfig
	commonconfig.LoadAppConfig(v, parser, &cfg)

	return &cfg
}

func loadEnv() {
	_ = godotenv.Load()
}
