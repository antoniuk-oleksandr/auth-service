package middleware

import (
	"time"

	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/presentation/http"
)

type RequestLogger interface {
	Handle() http.Handler
}

type requestLogger struct {
	lgr logger.Logger
}

func NewRequestLogger(lgr logger.Logger) RequestLogger {
	return &requestLogger{
		lgr: lgr,
	}
}

func (m *requestLogger) Handle() http.Handler {
	return func(ctx http.HTTPContext) error {
		start := time.Now()

		err := ctx.Next()

		m.lgr.Info("HTTP request",
			logger.NewStringField("method", ctx.Method()),
			logger.NewStringField("path", ctx.Path()),
			logger.NewIntField("status", ctx.StatusCode()),
			logger.NewFloatField("latency_ms", float64(time.Since(start).Milliseconds())),
		)

		return err
	}
}
