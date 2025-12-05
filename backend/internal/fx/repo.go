package fxmodules

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/db"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"

	"go.uber.org/fx"
)

var RepoModule = fx.Module("repo", fx.Provide(
	provideTxManager,
	provideUsersRepo,
))

func provideTxManager(factory db.RepositoryFactory) db.TxManager {
	return factory.CreateTxManager()
}

func provideUsersRepo(factory db.RepositoryFactory) users.Repository {
	return factory.CreateUsersRepo()
}
