package fiber

import (
	"context"

	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"

	fiberLib "github.com/gofiber/fiber/v2"
)

type FiberContext struct {
	ctx *fiberLib.Ctx
}

func NewFiberContext(c *fiberLib.Ctx) http.HTTPContext {
	return &FiberContext{ctx: c}
}

func (fc *FiberContext) BindJSON(obj any) error {
	return fc.ctx.BodyParser(obj)

}

func (fc *FiberContext) Context() context.Context {
	return fc.ctx.Context()
}

func (fc *FiberContext) JSON(obj any) error {
	return fc.ctx.JSON(obj)

}

func (fc *FiberContext) Param(key string) string {
	return fc.ctx.Params(key)
}

func (fc *FiberContext) Query(key string) string {
	return fc.ctx.Query(key)

}

func (fc *FiberContext) Status(code int) http.HTTPContext {
	fc.ctx.Status(code)
	return fc
}

func (fc *FiberContext) String(code int, s string) error {
	return fc.ctx.Status(code).SendString(s)
}

func (fc *FiberContext) Method() string {
	return fc.ctx.Method()
}

func (fc *FiberContext) Next() error {
	return fc.ctx.Next()
}

func (fc *FiberContext) Path() string {
	return fc.ctx.Path()
}

func (fc *FiberContext) StatusCode() int {
	return fc.ctx.Response().StatusCode()
}
