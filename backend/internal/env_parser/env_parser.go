package envparser

type EnvParser interface {
	Parse(v any) error
}
