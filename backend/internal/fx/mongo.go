package fxmodules

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/config"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/db"
	mongodb "github.com/antoniuk-oleksandr/auth-service/backend/internal/db/mongo"
	usersInfra "github.com/antoniuk-oleksandr/auth-service/backend/internal/infra/mongo/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"

	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/fx"
)

func provideMongoFactory(
	lc fx.Lifecycle,
	cfg *config.AppConfig,
	lgr logger.Logger,
	mapper usersInfra.Mapper,
) (db.RepositoryFactory, error) {
	client, err := mongodb.New(cfg.Mongo)
	if err != nil {
		return nil, err
	}

	registerMongoLifeCycle(lc, client, lgr)

	return mongodb.NewMongoFactory(client, mapper, cfg.Mongo.DBName, lgr), nil
}

func registerMongoLifeCycle(lc fx.Lifecycle, client *mongo.Client, l logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l.Info("Connecting to MongoDB")
			if err := client.Ping(ctx, nil); err != nil {
				return err
			}
			l.Info("Connected to MongoDB")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			l.Info("Disconnecting from MongoDB")
			return client.Disconnect(ctx)
		},
	})
}
