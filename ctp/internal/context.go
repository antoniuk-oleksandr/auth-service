package customprotoc

import (
	"context"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"

	"github.com/vmihailenco/msgpack/v5"
)

type ctx struct {
	req      *types.Request
	c        context.Context
	values   map[string]any
	handlers []types.HandlerFunc
	index    int
	status   types.Status
}

func NewCtx(req *types.Request, c context.Context, handlers []types.HandlerFunc) types.Ctx {
	return &ctx{
		req:      req,
		c:        c,
		values:   make(map[string]any),
		index:    0,
		handlers: handlers,
		status:   types.StatusOK,
	}
}

func (c *ctx) Status(status types.Status) types.Ctx {
	c.status = status
	return c
}

func (c *ctx) Send(data any) error {
	bodyBytes, err := msgpack.Marshal(data)
	if err != nil {
		return err
	}

	return &types.Response{
		Status: c.status,
		Body:   bodyBytes,
	}
}

func (c *ctx) Next() error {
	c.index++
	if c.index < len(c.handlers) {
		return c.handlers[c.index](c)
	}
	return nil
}

func (c *ctx) GetValue(key string) any {
	return c.values[key]
}

func (c *ctx) SetValue(key string, value any) {
	c.values[key] = value
}

func (c *ctx) Context() context.Context {
	return c.c
}

func (c *ctx) BodyParser(out any) error {
	return msgpack.Unmarshal(c.req.Body, out)
}
