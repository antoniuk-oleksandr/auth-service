package auth

import "errors"

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrFailedToLoginUser    = errors.New("failed to login user")
	ErrFailedToRegisterUser = errors.New("failed to register user")
)
