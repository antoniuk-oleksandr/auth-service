package fxmodules

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/config"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/db"
	usersMongoInfra "github.com/antoniuk-oleksandr/auth-service/backend/internal/infra/mongo/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"fmt"

	"go.uber.org/fx"
)

var DBModule = fx.Module("infrastructure",
	fx.Provide(provideRepoFactory),
)

func provideRepoFactory(
	lc fx.Lifecycle,
	cfg *config.AppConfig,
	l logger.Logger,
) (db.RepositoryFactory, error) {
	switch cfg.Database.Type {
	case "mongo":
		mapper := usersMongoInfra.NewMapper()
		return provideMongoFactory(lc, cfg, l, mapper)
	// For instance PostgreSQL can be added here
	// case "postgres":
	// 	return providePostgresFactory(lc, cfg, l)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}
}
