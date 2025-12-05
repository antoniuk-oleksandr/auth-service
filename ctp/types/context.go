package types

import "context"

type HandlerFunc func(ctx Ctx) error

type Map map[string]any

type Ctx interface {
	BodyParser(out any) error
	Context() context.Context
	SetValue(key string, value any)
	GetValue(key string) any
	Next() error
	Send(data any) error
	Status(status Status) Ctx
}
