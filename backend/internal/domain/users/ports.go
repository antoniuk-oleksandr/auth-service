package users

import (
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, cmd CreateUserCommand) (*User, error)
		FindByID(ctx context.Context, id string) (*User, error)
		FindByUsername(ctx context.Context, username string) (*User, error)
	}

	Service interface {
		CreateUser(ctx context.Context, cmd CreateUserCommand) (*User, error)
		GetUserByID(ctx context.Context, id string) (*User, error)
		GetUserByUsername(ctx context.Context, username string) (*User, error)
	}
)
