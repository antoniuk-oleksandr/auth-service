package auth

import "errors"

var (
	ErrFailedToRegisterUser = errors.New("failed to register user")
	ErrFailedToLoginUser = errors.New("failed to login user")
)
