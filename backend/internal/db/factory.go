package db

import (
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
)

type RepositoryFactory interface {
	CreateUsersRepo() usersDomain.Repository
	CreateTxManager() TxManager
}
