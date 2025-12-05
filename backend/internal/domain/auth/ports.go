package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Service interface {
		Register(ctx context.Context, cmd RegisterCommand) (*JWT, error)
		Login(ctx context.Context, cmd LoginCommand) (*JWT, error)
	}

	JWTManager interface {
		SignTokens(userID string, jti string) (dto *JWT, err error)
		SignAccessToken(userID, jti string) (string, error)
		SignRefreshToken(userID, jti string) (string, error)
		GenerateClaims(userID, jti string, ttl int) jwt.MapClaims
	}

	KeyGenerator interface {
		Generate() (string, error)
	}
)
