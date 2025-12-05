package logger

import loggeriface "github.com/antoniuk-oleksandr/auth-service/backend/internal/logger/logger_interface"

type (
	Field     = loggeriface.Field
	FieldType = loggeriface.FieldType
	Logger    = loggeriface.Logger
)

const (
	StringField = loggeriface.StringField
	IntField = loggeriface.IntField
	FloatField = loggeriface.FloatField
	BoolField = loggeriface.BoolField
	ErrorField = loggeriface.ErrorField
	AnyField = loggeriface.AnyField
)

func NewField(key string, val any, t FieldType) Field {
	return Field{Key: key, Type: t, Value: val}
}

func NewStringField(key, val string) Field {
	return Field{Key: key, Type: StringField, Value: val}
}

func NewIntField(key string, val int) Field {
	return Field{Key: key, Type: IntField, Value: val}
}

func NewBoolField(key string, val bool) Field {
	return Field{Key: key, Type: BoolField, Value: val}
}

func NewFloatField(key string, val float64) Field {
	return Field{Key: key, Type: FloatField, Value: val}
}

func NewErrField(err error) Field {
	return Field{Key: "error", Type: ErrorField, Value: err}
}

func NewAnyField(key string, val any) Field {
	return Field{Key: key, Type: AnyField, Value: val}
}
