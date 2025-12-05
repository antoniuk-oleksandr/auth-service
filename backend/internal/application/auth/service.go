package auth

import (
	"context"
	"errors"

	authMapper "github.com/antoniuk-oleksandr/auth-service/backend/internal/application/auth/mapper"
	authDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"github.com/antoniuk-oleksandr/auth-service/backend/pkg/hasher"
	uuidgeneator "github.com/antoniuk-oleksandr/auth-service/backend/pkg/uuid"
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
	s.lgr.Info("Login attempt",
		logger.NewStringField("username", cmd.Username),
	)

	user, err := s.usersService.GetUserByUsername(ctx, cmd.Username)
	if err != nil {
		if errors.Is(err, usersDomain.ErrUserNotFound) {
			s.lgr.Warn("Invalid credentials: user not found",
				logger.NewStringField("username", cmd.Username),
			)
			return nil, authDomain.ErrInvalidCredentials
		}
		return nil, err
	}

	if err := s.hasher.Compare(cmd.Password, user.PasswordHash); err != nil {
		s.lgr.Warn("Invalid credentials: wrong password",
			logger.NewStringField("username", cmd.Username),
		)
		return nil, authDomain.ErrInvalidCredentials
	}

	jti := s.uuidGenerator.Generate()
	jwt, err := s.jwtManager.SignTokens(user.ID, jti)
	if err != nil {
		s.lgr.Error("Failed to sign JWT tokens",
			logger.NewStringField("username", cmd.Username),
			logger.NewErrField(err),
		)
		return nil, authDomain.ErrInvalidCredentials
	}

	s.lgr.Info("Login successful",
		logger.NewStringField("user_id", user.ID),
		logger.NewStringField("username", cmd.Username),
	)

	return jwt, nil
}

func (s *service) Register(ctx context.Context, cmd authDomain.RegisterCommand) (*authDomain.JWT, error) {
	s.lgr.Info("User registration attempt",
		logger.NewStringField("username", cmd.Username),
	)

	hashedPassword, err := s.hasher.Hash(cmd.Password)
	if err != nil {
		s.lgr.Error("Failed to hash password",
			logger.NewStringField("username", cmd.Username),
			logger.NewErrField(err),
		)
		return nil, authDomain.ErrInvalidCredentials
	}

	userDto := s.mapper.ToCreateUserCommand(cmd, hashedPassword)

	createdUser, err := s.usersService.CreateUser(ctx, userDto)
	if err != nil {
		// repo already logs DB errors
		if errors.Is(err, usersDomain.ErrUsernameTaken) {
			s.lgr.Warn("Username already taken",
				logger.NewStringField("username", cmd.Username),
			)
			return nil, err
		}

		return nil, err
	}

	jti := s.uuidGenerator.Generate()
	jwt, err := s.jwtManager.SignTokens(createdUser.ID, jti)
	if err != nil {
		s.lgr.Error("Failed to sign JWT tokens",
			logger.NewStringField("username", cmd.Username),
			logger.NewErrField(err),
		)
		return nil, authDomain.ErrInvalidCredentials
	}

	s.lgr.Info("User registered successfully",
		logger.NewStringField("user_id", createdUser.ID),
		logger.NewStringField("username", createdUser.Username),
	)

	return jwt, nil
}
