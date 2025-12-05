package mongodb

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/db"
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	usersInfra "github.com/antoniuk-oleksandr/auth-service/backend/internal/infra/mongo/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type mongoFactory struct {
	client *mongo.Client
	db     *mongo.Database
	mapper usersInfra.Mapper
	lgr    logger.Logger
}

func NewMongoFactory(
	client *mongo.Client,
	mapper usersInfra.Mapper,
	dbName string,
	lgr logger.Logger,
) db.RepositoryFactory {
	return &mongoFactory{
		client: client,
		mapper: mapper,
		db:     client.Database(dbName),
		lgr:    lgr,
	}
}

func (m *mongoFactory) CreateTxManager() db.TxManager {
	return NewTransactionManager(m.client)
}

func (m *mongoFactory) CreateUsersRepo() usersDomain.Repository {
	return usersInfra.NewRepository(m.db, m.mapper, m.lgr)
}
