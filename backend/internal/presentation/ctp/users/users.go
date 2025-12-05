package users

import (
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/ctp/common"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
)

type controller struct {
	srv usersDomain.Service
}

type Controller interface {
	GetUserByID(ctx types.Ctx) error
}

func NewController(srv usersDomain.Service) Controller {
	return &controller{
		srv: srv,
	}
}

func (ctrl *controller) GetUserByID(ctx types.Ctx) error {
	var id string
	if err := common.ParseBody(ctx, &id); err != nil {
		return err
	}

	user, err := ctrl.srv.GetUserByID(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(types.StatusOK).Send(user)
}
