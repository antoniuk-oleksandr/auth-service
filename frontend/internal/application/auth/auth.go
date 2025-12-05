package auth

import (
	authDomain "github.com/antoniuk-oleksandr/auth-service/frontend/internal/domain/auth"
	usersDomain "github.com/antoniuk-oleksandr/auth-service/frontend/internal/domain/users"

	"github.com/antoniuk-oleksandr/auth-service/ctp/client"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
)

type authService struct {
	client client.Client
}

func NewService(client client.Client) authDomain.AuthService {
	return &authService{
		client: client,
	}
}

func (srv *authService) GetUserProfile(token string) (*usersDomain.User, error) {
	panic("unimplemented")
}

func (srv *authService) LoginUser(
	credentials *authDomain.Credentials,
) (*authDomain.JWT, error) {
	var jwt *authDomain.JWT
	resp, err := srv.client.Send("auth.login", credentials, &jwt)
	if err != nil {
		return nil, authDomain.ErrFailedToLoginUser
	}

	switch resp.Status {
	case types.StatusCreated:
		return jwt, nil
	case types.StatusUnauthorized:
		return nil, authDomain.ErrInvalidCredentials
	case types.StatusBadRequest:
		return nil, authDomain.ErrInvalidCredentials
	case types.StatusNotFound:
		return nil, authDomain.ErrInvalidCredentials
	case types.StatusInternalError:
		return nil, authDomain.ErrFailedToLoginUser
	default:
		return nil, authDomain.ErrFailedToLoginUser
	}
}

func (srv *authService) RegisterUser(
	credentials *authDomain.Credentials,
) (*authDomain.JWT, error) {
	var jwt *authDomain.JWT
	resp, err := srv.client.Send("auth.register", credentials, &jwt)
	if err != nil {
		return nil, authDomain.ErrFailedToRegisterUser
	}

	if resp.Status != types.StatusCreated {
		switch resp.Status {
		case types.StatusCreated:
			return jwt, nil
		case types.StatusUnauthorized:
			return nil, authDomain.ErrInvalidCredentials
		case types.StatusBadRequest:
			return nil, authDomain.ErrInvalidCredentials
		case types.StatusNotFound:
			return nil, authDomain.ErrInvalidCredentials
		case types.StatusInternalError:
			return nil, authDomain.ErrFailedToRegisterUser
		default:
			return nil, authDomain.ErrFailedToRegisterUser
		}
	}

	return jwt, nil
}