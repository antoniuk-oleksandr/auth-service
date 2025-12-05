package types

import "io"

type Response struct {
	Status Status
	Body   []byte
}

func (r Response) Error() string {
	return "response error"
}

type ResponseWriter interface {
	WriteResponse(w io.Writer, resp *Response) error
}
