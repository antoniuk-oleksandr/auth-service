package types

type Field struct {
	Key   string
	Value any
	Type  FieldType
}

type FieldType int

const (
	StringField FieldType = iota
	IntField
	FloatField
	BoolField
	ErrorField
	AnyField
)

type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
}
