package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrAuthFailed         = errors.New("authentication failed")
)
