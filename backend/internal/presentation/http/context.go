package http

import "context"

type HTTPContext interface {
	Param(key string) string
	Query(key string) string
	JSON(obj any) error
	String(code int, s string) error
	BindJSON(obj any) error
	Status(code int) HTTPContext
	Context() context.Context
	Method() string
	Path() string
	StatusCode() int
	Next() error
}
