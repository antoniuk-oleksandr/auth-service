package fxmodules

import (
	authApp "github.com/antoniuk-oleksandr/auth-service/backend/internal/application/auth"
	authMapper "github.com/antoniuk-oleksandr/auth-service/backend/internal/application/auth/mapper"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/config"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/db"
	authDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"github.com/antoniuk-oleksandr/auth-service/backend/pkg/hasher"
	uuidgeneator "github.com/antoniuk-oleksandr/auth-service/backend/pkg/uuid"

	"go.uber.org/fx"
)

var AuthAppModule = fx.Module("auth",
	fx.Provide(authMapper.New),
	fx.Provide(provideJWTManager),
	fx.Provide(provideAuthServiceWithTx),
)

func provideJWTManager(cfg *config.AppConfig) authDomain.JWTManager {
	return authApp.NewJWTManager(
		cfg.JWT.SecretKey,
		cfg.JWT.AccessTokenTTL,
		cfg.JWT.RefreshTokenTTL,
	)
}

func provideAuthServiceWithTx(
	usersSrv usersDomain.Service,
	jwtManager authDomain.JWTManager,
	hasher hasher.Hasher,
	uuidGen uuidgeneator.UUIDGenerator,
	mapper authMapper.Mapper,
	factory db.RepositoryFactory,
	lgr logger.Logger,
) authDomain.Service {
	authSrv := authApp.NewService(usersSrv, jwtManager, hasher, uuidGen, mapper, lgr)
	return authApp.NewServiceWithTx(authSrv, factory.CreateTxManager())
}
