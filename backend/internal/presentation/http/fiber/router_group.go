package fiber

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"

	fiberlib "github.com/gofiber/fiber/v2"
)

type FiberRouterGroup struct {
	group fiberlib.Router
}

func (fg *FiberRouterGroup) Get(path string, handler http.Handler) http.RouterGroup {
	fg.group.Get(path, fg.adaptHandler(handler))
	return fg
}

func (fg *FiberRouterGroup) Post(path string, handler http.Handler) http.RouterGroup {
	fg.group.Post(path, fg.adaptHandler(handler))
	return fg
}

func (fg *FiberRouterGroup) Put(path string, handler http.Handler) http.RouterGroup {
	fg.group.Put(path, fg.adaptHandler(handler))
	return fg
}

func (fg *FiberRouterGroup) Delete(path string, handler http.Handler) http.RouterGroup {
	fg.group.Delete(path, fg.adaptHandler(handler))
	return fg
}