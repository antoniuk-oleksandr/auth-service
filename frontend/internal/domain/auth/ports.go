package auth

import usersDomain "github.com/antoniuk-oleksandr/auth-service/frontend/internal/domain/users"

type AuthService interface {
	LoginUser(credentials *Credentials) (*JWT, error)
	RegisterUser(credentials *Credentials) (*JWT, error)
	GetUserProfile(token string) (*usersDomain.User, error)
}
