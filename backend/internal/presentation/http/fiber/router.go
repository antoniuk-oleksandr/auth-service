package fiber

import (
	"fmt"

	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"

	fiberLib "github.com/gofiber/fiber/v2"
)

type FiberRouter struct {
	app *fiberLib.App
}

func NewFiberRouter(config any) (http.HTTPRouter, error) {
	cfg, ok := config.(fiberLib.Config)
	if !ok {
		return nil, fmt.Errorf("invalid config type for FiberRouter")
	}

	return &FiberRouter{
		app: fiberLib.New(cfg),
	}, nil
}

func (f *FiberRouter) Use(middleware any) {
	f.app.Use(middleware)
}

func (f *FiberRouter) Start(address string) error {
	return f.app.Listen(address)
}

func (f *FiberRouter) Group(path string) http.RouterGroup {
	return &FiberRouterGroup{group: f.app.Group(path)}
}

func (f *FiberRouter) Close() error {
	return f.app.Shutdown()
}
