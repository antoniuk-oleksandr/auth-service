package auth

import (
	"time"

	ports "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"

	"github.com/golang-jwt/jwt/v5"
)

type jwtHelper struct {
	secret     []byte
	accessTTL  int
	refreshTTL int
}

func NewJWTManager(secret string, accessTTL, refreshTTL int) ports.JWTManager {
	return &jwtHelper{
		secret:     []byte(secret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (j *jwtHelper) GenerateClaims(userID string, jti string, ttl int) jwt.MapClaims {
	return jwt.MapClaims{
		"sub": userID,
		"jti": jti,
		"exp": time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
		"iat": time.Now().Unix(),
	}
}

func (j *jwtHelper) SignAccessToken(userID string, jti string) (string, error) {
	claims := j.GenerateClaims(userID, jti, j.accessTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *jwtHelper) SignRefreshToken(userID string, jti string) (string, error) {
	claims := j.GenerateClaims(userID, jti, j.refreshTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *jwtHelper) SignTokens(userID string, jti string) (dto *ports.JWT, err error) {
	accessToken, err := j.SignAccessToken(userID, jti)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.SignRefreshToken(userID, jti)
	if err != nil {
		return nil, err
	}

	return &ports.JWT{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  j.accessTTL,
		RefreshExpiresIn: j.refreshTTL,
	}, nil
}
