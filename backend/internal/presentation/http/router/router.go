package http

import (
	"fmt"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http/fiber"
)

var factories = map[string]func(config any) (http.HTTPRouter, error){
	"fiber": fiber.NewFiberRouter,
}

func NewHTTPRouter(factory string, config any) (http.HTTPRouter, error) {
	f, ok := factories[factory]
	if !ok {
		return nil, fmt.Errorf("invalid router factory: %s", factory)
	}
	return f(config)
}
