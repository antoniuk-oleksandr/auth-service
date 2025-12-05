package validator

type Validator interface {
	Struct(val any) error
}