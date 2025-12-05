package zaplogger

import (
	loggeriface "github.com/antoniuk-oleksandr/auth-service/backend/internal/logger/logger_interface"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	l *zap.Logger
}

func NewZapLogger(core zapcore.Core) (loggeriface.Logger, error) {
	l, err := zap.NewProduction(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		return nil, err
	}

	return &zapLogger{
		l: l,
	}, nil
}

func (z *zapLogger) Error(msg string, fields ...loggeriface.Field) {
	zapFields := convertToZapFields(fields)
	z.l.Error(msg, zapFields...)
}

func (z *zapLogger) Fatal(msg string, fields ...loggeriface.Field) {
	zapFields := convertToZapFields(fields)
	z.l.Fatal(msg, zapFields...)
}

func (z *zapLogger) Info(msg string, fields ...loggeriface.Field) {
	zapFields := convertToZapFields(fields)
	z.l.Info(msg, zapFields...)
}

func (z *zapLogger) Warn(msg string, fields ...loggeriface.Field) {
	zapFields := convertToZapFields(fields)
	z.l.Warn(msg, zapFields...)
}

func (z *zapLogger) Debug(msg string, fields ...loggeriface.Field) {
	zapFields := convertToZapFields(fields)
	z.l.Debug(msg, zapFields...)
}

func convertToZapFields(fields []loggeriface.Field) []zap.Field {
	var zapFields []zap.Field
	for _, f := range fields {
		switch f.Type {
		case loggeriface.StringField:
			zapFields = append(zapFields, zap.String(f.Key, f.Value.(string)))
		case loggeriface.IntField:
			zapFields = append(zapFields, zap.Int(f.Key, f.Value.(int)))
		case loggeriface.BoolField:
			zapFields = append(zapFields, zap.Bool(f.Key, f.Value.(bool)))
		case loggeriface.FloatField:
			zapFields = append(zapFields, zap.Float64(f.Key, f.Value.(float64)))
		case loggeriface.ErrorField:
			zapFields = append(zapFields, zap.Error(f.Value.(error)))
		case loggeriface.AnyField:
			zapFields = append(zapFields, zap.Any(f.Key, f.Value))
		}
	}

	return zapFields
}
