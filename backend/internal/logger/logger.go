package logger

import loggeriface "github.com/antoniuk-oleksandr/auth-service/backend/internal/logger/logger_interface"

type (
	Field     = loggeriface.Field
	FieldType = loggeriface.FieldType
	Logger    = loggeriface.Logger
)

const (
	StringField FieldType = iota
	IntField
	FloatField
	BoolField
	ErrorField
	AnyField
)

func String(key, val string) Field {
	return Field{Key: key, Type: StringField, Value: val}
}

func Int(key string, val int) Field {
	return Field{Key: key, Type: IntField, Value: val}
}

func Bool(key string, val bool) Field {
	return Field{Key: key, Type: BoolField, Value: val}
}

func Float(key string, val float64) Field {
	return Field{Key: key, Type: FloatField, Value: val}
}

func Err(err error) Field {
	return Field{Key: "error", Type: ErrorField, Value: err}
}
