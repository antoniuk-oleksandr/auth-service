package customprotoc

import (
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
	"fmt"
	"io"
)

type Responses = types.Response



type defaultResponseWriter struct{}

func NewResponseWriter() types.ResponseWriter {
	return &defaultResponseWriter{}
}

func (w *defaultResponseWriter) WriteResponse(writer io.Writer, resp *types.Response) error {
	encoded, err := EncodeResponse(resp)
	if err != nil {
		return fmt.Errorf("encode response: %w", err)
	}

	if _, err := writer.Write(encoded); err != nil {
		return fmt.Errorf("write response: %w", err)
	}

	return nil
}
