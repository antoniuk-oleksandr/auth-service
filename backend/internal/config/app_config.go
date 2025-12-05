package config

import (
	envparser "github.com/antoniuk-oleksandr/auth-service/backend/internal/env_parser"
	val "github.com/antoniuk-oleksandr/auth-service/backend/pkg/validator"
)

type MongoConfig struct {
	DBName string `env:"MONGO_DBNAME" validate:"required"`
	URI    string `env:"MONGO_URI" validate:"required,url"`
}

type HasherConfig struct {
	Cost int `env:"HASHER_COST" validate:"gte=4,lte=31"`
}

type JWTConfig struct {
	SecretKey       string `env:"JWT_SECRET_KEY" validate:"required"`
	AccessTokenTTL  int    `env:"JWT_ACCESS_TOKEN_TTL" validate:"gte=0"`
	RefreshTokenTTL int    `env:"JWT_REFRESH_TOKEN_TTL" validate:"gte=0"`
}

type ServerConfig struct {
	Port          string `env:"SERVER_PORT" validate:"required"`
	HTTPFramework string `env:"HTTP_FRAMEWORK" validate:"required,oneof=fiber gin"`
}

type LoggerConfig struct {
	Type string `env:"LOGGER_TYPE" validate:"required,oneof=zap_console zap_json"`
}

type DatabaseConfig struct {
	Type string `env:"DATABASE_TYPE" validate:"required,oneof=mongo postgres"`
}

type AppConfig struct {
	Mongo    MongoConfig
	Server   ServerConfig
	JWT      JWTConfig
	Hasher   HasherConfig
	Logger   LoggerConfig
	Database DatabaseConfig
}

func LoadAppConfig(v val.Validator, parser envparser.EnvParser) (*AppConfig, error) {
	var cfg AppConfig

	if err := parser.Parse(&cfg); err != nil {
		return nil, err
	}

	if err := v.Struct(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
