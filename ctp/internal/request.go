package customprotoc

import (
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
	"encoding/binary"
	"fmt"
	"io"
)

type defaultRequestReader struct {
	maxBodySize uint32
}

func NewRequestReader(maxBodySize uint32) types.RequestReader {
	return &defaultRequestReader{
		maxBodySize: maxBodySize,
	}
}

func (r *defaultRequestReader) ReadRequest(reader io.Reader) (*types.Request, error) {
	var commandLen uint16
	if err := binary.Read(reader, binary.BigEndian, &commandLen); err != nil {
		return nil, fmt.Errorf("read command length: %w", err)
	}

	if commandLen > MaxCommandSize {
		return nil, ErrCommandTooLong
	}

	cmdBytes := make([]byte, commandLen)
	if _, err := io.ReadFull(reader, cmdBytes); err != nil {
		return nil, fmt.Errorf("read command: %w", err)
	}

	var bodyLen uint32
	if err := binary.Read(reader, binary.BigEndian, &bodyLen); err != nil {
		return nil, fmt.Errorf("read body length: %w", err)
	}

	if bodyLen > r.maxBodySize {
		return nil, ErrBodyTooLarge
	}

	body := make([]byte, bodyLen)
	if bodyLen > 0 {
		if _, err := io.ReadFull(reader, body); err != nil {
			return nil, fmt.Errorf("read body: %w", err)
		}
	}

	return &types.Request{
		Command: string(cmdBytes),
		Body:    body,
	}, nil
}
