package auth

import (
	authDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
)

type mapper struct{}

type Mapper interface {
	ToCreateUserCommand(cmd authDomain.RegisterCommand, hashedPassword string) usersDomain.CreateUserCommand
}

func New() Mapper {
	return &mapper{}
}

func (m *mapper) ToCreateUserCommand(
	cmd authDomain.RegisterCommand,
	hashedPassword string,
) usersDomain.CreateUserCommand {
	return usersDomain.CreateUserCommand{
		Username:     cmd.Username,
		PasswordHash: hashedPassword,
	}
}
