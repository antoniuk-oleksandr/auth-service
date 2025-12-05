package customprotoc

import (
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
	"fmt"
)

type DefaultLogger struct{}

func NewDefaultLogger() types.Logger {
	return &DefaultLogger{}
}

func (d *DefaultLogger) Debug(msg string, fields ...types.Field) {
	d.print("DEBUG", msg, fields...)
}

func (d *DefaultLogger) Info(msg string, fields ...types.Field) {
	d.print("INFO", msg, fields...)
}

func (d *DefaultLogger) Warn(msg string, fields ...types.Field) {
	d.print("WARN", msg, fields...)
}

func (d *DefaultLogger) Error(msg string, fields ...types.Field) {
	d.print("ERROR", msg, fields...)
}

func (d *DefaultLogger) Fatal(msg string, fields ...types.Field) {
	d.print("FATAL", msg, fields...)
}

func (d *DefaultLogger) print(level string, msg string, fields ...types.Field) {
	fmt.Printf("[%s] %s", level, msg)
	if len(fields) > 0 {
		fmt.Printf(" | ")
		for i, f := range fields {
			if i > 0 {
				fmt.Print(", ")
			}
			switch f.Type {
			case types.ErrorField:
				if errVal, ok := f.Value.(error); ok {
					fmt.Printf("%s=%v", f.Key, errVal.Error())
				} else {
					fmt.Printf("%s=%v", f.Key, f.Value)
				}
			default:
				fmt.Printf("%s=%v", f.Key, f.Value)
			}
		}
	}
	fmt.Println()
}
