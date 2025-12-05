package users

import (
	"errors"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUsernameTaken      = errors.New("username is already taken")
	ErrFailedToCreateUser = errors.New("failed to create user")
	ErrFailedToFindUser   = errors.New("failed to find user")
)
