package users

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/common/registry"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
)

type routes struct {
	conn types.Server
	ctrl Controller
}

func NewRoutes(conn types.Server, ctrl Controller) registry.Routes {
	return &routes{
		conn: conn,
		ctrl: ctrl,
	}
}

func (r *routes) Register() {
	r.conn.RegisterHandler("users.getById", r.ctrl.GetUserByID)
}
