package commonvalidator

type Validator interface {
	Struct(val any) error
}