package users

import (
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"context"
)

type service struct {
	repo usersDomain.Repository
	lgr  logger.Logger
}

func NewService(repo usersDomain.Repository, lgr logger.Logger) usersDomain.Service {
	return &service{
		repo: repo,
		lgr:  lgr,
	}
}

func (s *service) GetUserByUsername(ctx context.Context, username string) (*usersDomain.User, error) {
	return s.repo.FindByUsername(ctx, username)
}

func (srv *service) CreateUser(ctx context.Context, user usersDomain.CreateUserCommand) (*usersDomain.User, error) {
	createdUser, err := srv.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*usersDomain.User, error) {
	return s.repo.FindByID(ctx, id)
}