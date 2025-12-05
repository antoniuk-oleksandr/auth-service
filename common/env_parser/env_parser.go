package commonenvparser

type EnvParser interface {
	Parse(v any) error
}
