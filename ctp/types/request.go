package types

import (
	"io"
)

type Request struct {
	Command string
	Body    []byte
}

type RequestReader interface {
	ReadRequest(r io.Reader) (*Request, error)
}
