package client

import (
	"encoding/binary"
	"net"

	customprotoc "github.com/antoniuk-oleksandr/auth-service/ctp/internal"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"

	"github.com/vmihailenco/msgpack/v5"
)

type Client interface {
	Send(command string, reqBody any, respBodyOut any) (*types.Response, error)
}

type client struct {
	addr string
}

func NewClient(addr string) Client {
	return &client{addr: addr}
}

func (c *client) Send(command string, reqBody any, respBodyOut any) (*types.Response, error) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var bodyBytes []byte
	if reqBody != nil {
		bodyBytes, err = msgpack.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
	}

	req := &types.Request{
		Command: command,
		Body:    bodyBytes,
	}

	reqBytes, err := customprotoc.EncodeRequest(req)
	if err != nil {
		return nil, err
	}

	if _, err := conn.Write(reqBytes); err != nil {
		return nil, err
	}

	respHeader := make([]byte, 8)
	if _, err := conn.Read(respHeader); err != nil {
		return nil, err
	}

	respBodyLen := binary.BigEndian.Uint32(respHeader[4:8])
	if respBodyLen > customprotoc.MaxBodySize {
		return nil, customprotoc.ErrBodyTooLarge
	}

	respBody := make([]byte, respBodyLen)
	if _, err := conn.Read(respBody); err != nil {
		return nil, err
	}

	resp := &types.Response{
		Status: types.Status(binary.BigEndian.Uint32(respHeader[0:4])),
		Body:   respBody,
	}

	if respBodyOut != nil {
		if err := msgpack.Unmarshal(resp.Body, respBodyOut); err != nil {
			return nil, err
		}
	}

	return resp, nil
}
