package customprotoc

import (
	"bytes"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
	"encoding/binary"

	"github.com/vmihailenco/msgpack/v5"
)

func EncodeRequest(req *types.Request) ([]byte, error) {
	if len(req.Command) > int(MaxCommandSize) {
		return nil, ErrCommandTooLong
	}

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, uint16(len(req.Command))); err != nil {
		return nil, err
	}

	buf.WriteString(req.Command)

	if err := binary.Write(buf, binary.BigEndian, uint32(len(req.Body))); err != nil {
		return nil, err
	}

	buf.Write(req.Body)

	return buf.Bytes(), nil
}

func DecodeRequest(data []byte, outBody any) (*types.Request, error) {
	if len(data) < 2+4 {
		return nil, ErrInvalidHeader
	}

	buf := bytes.NewReader(data)

	var commandLen uint16
	if err := binary.Read(buf, binary.BigEndian, &commandLen); err != nil {
		return nil, err
	}

	if commandLen > MaxCommandSize {
		return nil, ErrCommandTooLong
	}

	command := make([]byte, commandLen)
	if _, err := buf.Read(command); err != nil {
		return nil, err
	}

	var bodyLen uint32
	if err := binary.Read(buf, binary.BigEndian, &bodyLen); err != nil {
		return nil, err
	}

	if bodyLen > MaxBodySize {
		return nil, ErrBodyTooLarge
	}

	bodyBytes := make([]byte, bodyLen)
	if _, err := buf.Read(bodyBytes); err != nil {
		return nil, err
	}

	if outBody != nil {
		if err := msgpack.Unmarshal(bodyBytes, outBody); err != nil {
			return nil, err
		}
	}

	return &types.Request{
		Command: string(command),
		Body:    bodyBytes,
	}, nil
}

func EncodeResponse(resp *types.Response) ([]byte, error) {
	if len(resp.Body) > int(MaxBodySize) {
		return nil, ErrBodyTooLarge
	}

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, int32(resp.Status)); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, uint32(len(resp.Body))); err != nil {
		return nil, err
	}

	buf.Write(resp.Body)

	return buf.Bytes(), nil
}

func DecodeResponse(data []byte, outBody any) (*types.Response, error) {
	if len(data) < 8 {
		return nil, ErrInvalidHeader
	}

	buf := bytes.NewReader(data)

	var status int32
	if err := binary.Read(buf, binary.BigEndian, &status); err != nil {
		return nil, err
	}

	var bodyLen uint32
	if err := binary.Read(buf, binary.BigEndian, &bodyLen); err != nil {
		return nil, err
	}

	bodyBytes := make([]byte, bodyLen)
	if _, err := buf.Read(bodyBytes); err != nil {
		return nil, err
	}

	if outBody != nil {
		if err := msgpack.Unmarshal(bodyBytes, outBody); err != nil {
			return nil, err
		}
	}

	return &types.Response{
		Status: types.Status(status),
		Body:   bodyBytes,
	}, nil
}
