package logger

import (
	"fmt"
	zapConsoleLogger "github.com/antoniuk-oleksandr/auth-service/backend/internal/logger/zap_console"
	zapJSONLogger "github.com/antoniuk-oleksandr/auth-service/backend/internal/logger/zap_json"
)

func New(t string) (Logger, error) {
	switch t {
	case "zap_json":
		return zapJSONLogger.NewZapJSONLogger()
	case "zap_console":
		return zapConsoleLogger.NewZapConsoleLogger()
	default:
		return nil, fmt.Errorf("unsupported logger type: %v", t)
	}
}
