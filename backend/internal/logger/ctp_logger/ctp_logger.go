package ctplogger

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
)

type ctpLogger struct {
	lgr logger.Logger
}

func NewCTPLogger(lgr logger.Logger) types.Logger {
	return &ctpLogger{
		lgr: lgr,
	}
}

func (c *ctpLogger) Debug(msg string, fields ...types.Field) {
	normalFields := c.convertFields(fields...)
	c.lgr.Debug(msg, normalFields...)
}

func (c *ctpLogger) Error(msg string, fields ...types.Field) {
	normalFields := c.convertFields(fields...)
	c.lgr.Error(msg, normalFields...)
}

func (c *ctpLogger) Fatal(msg string, fields ...types.Field) {
	normalFields := c.convertFields(fields...)
	c.lgr.Fatal(msg, normalFields...)
}

func (c *ctpLogger) Info(msg string, fields ...types.Field) {
	normalFields := c.convertFields(fields...)
	c.lgr.Info(msg, normalFields...)
}

func (c *ctpLogger) Warn(msg string, fields ...types.Field) {
	normalFields := c.convertFields(fields...)
	c.lgr.Warn(msg, normalFields...)
}

func (c *ctpLogger) convertFields(fields ...types.Field) []logger.Field {
	var loggerFields []logger.Field
	for _, f := range fields {
		normalField := logger.Field{
			Key:   f.Key,
			Value: f.Value,
			Type:  logger.FieldType(f.Type),
		}
		loggerFields = append(loggerFields, normalField)
	}
	return loggerFields
}
