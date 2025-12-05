package common

import "github.com/antoniuk-oleksandr/auth-service/ctp/types"

func ParseBody(ctx types.Ctx, out any) error {
	if err := ctx.BodyParser(&out); err != nil {
		_ = ctx.Status(types.StatusBadRequest).Send(types.Map{
			"error": "invalid request body",
		})
		return err
	}
	return nil
}
