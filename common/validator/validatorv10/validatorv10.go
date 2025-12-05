package commonvalidatorv10

import (
	commonvalidator "auth-common/validator"

	val "github.com/go-playground/validator/v10"
)

type validatorV10 struct {
	validate *val.Validate
}

func New() commonvalidator.Validator {
	return &validatorV10{
		validate: val.New(),
	}
}

func (v *validatorV10) Struct(val any) error {
	return v.validate.Struct(val)
}
