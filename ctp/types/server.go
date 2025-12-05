package types

import (
	"io"
)

type Server interface {
	RegisterHandler(command string, handlers ...HandlerFunc)
	Start(port string) error
	HandleConnection(conn io.ReadWriteCloser)
	ProcessRequest(req *Request) *Response
}
