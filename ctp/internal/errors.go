package customprotoc

import "errors"

var (
	ErrInvalidHeader   = errors.New("invalid header")
	ErrCommandTooLong  = errors.New("command too long")
	ErrBodyTooLarge    = errors.New("body too large")
	ErrHandlerNotFound = errors.New("handler not found")
	ErrInvalidResponse = errors.New("invalid response")
)

const (
	MaxCommandSize uint16 = 256
	MaxBodySize    uint32 = 1 * 1024 * 1024 // 1 MB
)
