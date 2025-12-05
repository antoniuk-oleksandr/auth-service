package zapconsolelogger

import (
	"os"
	loggeriface "github.com/antoniuk-oleksandr/auth-service/backend/internal/logger/logger_interface"
	zaplogger "github.com/antoniuk-oleksandr/auth-service/backend/internal/logger/zap"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapJSONLogger() (loggeriface.Logger, error) {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zaplogger.NewZapLogger(core)
}
