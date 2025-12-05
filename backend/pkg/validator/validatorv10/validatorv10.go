package validatorv10

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/pkg/validator"

	val "github.com/go-playground/validator/v10"
)

type validatorV10 struct {
	validate *val.Validate
}

func New() validator.Validator {
	return &validatorV10{
		validate: val.New(),
	}
}

func (v *validatorV10) Struct(val any) error {
	return v.validate.Struct(val)
}
