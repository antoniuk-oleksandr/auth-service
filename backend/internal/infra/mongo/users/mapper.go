package users

import (
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
)

	type Mapper interface {
	MapUserModelToDomain(model User) usersDomain.User
}

type mapper struct {
}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) MapUserModelToDomain(model User) usersDomain.User {
	return usersDomain.User{
		Username:     model.Username,
		PasswordHash: model.PasswordHash,
	}
}
