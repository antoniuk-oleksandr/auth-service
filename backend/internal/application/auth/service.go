package auth

import (
	authMapper "github.com/antoniuk-oleksandr/auth-service/backend/internal/application/auth/mapper"
	authDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"github.com/antoniuk-oleksandr/auth-service/backend/pkg/hasher"
	uuidgeneator "github.com/antoniuk-oleksandr/auth-service/backend/pkg/uuid"
	"context"
	"errors"
)

type service struct {
	usersService  usersDomain.Service
	jwtManager    authDomain.JWTManager
	hasher        hasher.Hasher
	uuidGenerator uuidgeneator.UUIDGenerator
	mapper        authMapper.Mapper
	lgr           logger.Logger
}

func NewService(
	usersService usersDomain.Service,
	jwtManager authDomain.JWTManager,
	hasher hasher.Hasher,
	uuidGenerator uuidgeneator.UUIDGenerator,
	mapper authMapper.Mapper,
	lgr logger.Logger,
) authDomain.Service {
	return &service{
		usersService:  usersService,
		jwtManager:    jwtManager,
		hasher:        hasher,
		uuidGenerator: uuidGenerator,
		mapper:        mapper,
		lgr:           lgr,
	}
}

func (s *service) Login(ctx context.Context, cmd authDomain.LoginCommand) (*authDomain.JWT, error) {
	user, err := s.usersService.GetUserByUsername(ctx, cmd.Username)
	if err != nil {
		if errors.Is(err, usersDomain.ErrUserNotFound) {
			return nil, authDomain.ErrInvalidCredentials
		}

		return nil, err
	}

	err = s.hasher.Compare(cmd.Password, user.PasswordHash)
	if err != nil {
		return nil, authDomain.ErrInvalidCredentials
	}

	jti := s.uuidGenerator.Generate()

	return s.jwtManager.SignTokens(user.ID, jti)
}

func (s *service) Register(ctx context.Context, cmd authDomain.RegisterCommand) (*authDomain.JWT, error) {
	hashedPassword, err := s.hasher.Hash(cmd.Password)
	if err != nil {
		return nil, authDomain.ErrInvalidCredentials
	}

	userDto := s.mapper.ToCreateUserCommand(cmd, hashedPassword)

	createdUser, err := s.usersService.CreateUser(ctx, userDto)
	if err != nil {
		return nil, err
	}

	jti := s.uuidGenerator.Generate()

	jwt, err := s.jwtManager.SignTokens(createdUser.ID, jti)
	if err != nil {
		return nil, authDomain.ErrInvalidCredentials
	}

	return jwt, err
}
