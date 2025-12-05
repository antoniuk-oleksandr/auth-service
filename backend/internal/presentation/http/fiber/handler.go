package fiber

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"

	fiberlib "github.com/gofiber/fiber/v2"
)

func (fg *FiberRouterGroup) adaptHandler(h http.Handler) fiberlib.Handler {
	return func(c *fiberlib.Ctx) error {
		ctx := NewFiberContext(c)
		return h(ctx)
	}
}
