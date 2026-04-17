package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func ProvideLogger(env string, serviceName string) *Logger {
	var config zap.Config

	if env == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	config.OutputPaths = []string{"stdout"}

	log, _ := config.Build(zap.Fields(
		zap.String("service", serviceName),
	))

	return &Logger{log}
}

type RequestData struct {
	Function  string      `json:"function"`
	ProcessID string      `json:"process_id"`
	IPAddress string      `json:"ip_address"`
	Request   interface{} `json:"request"`
}

func (l *Logger) LogRequest(ctx context.Context, data RequestData) {
	l.Info("Incoming Request",
		zap.String("func", data.Function),
		zap.String("p_id", data.ProcessID),
		zap.String("ip", data.IPAddress),
		zap.Any("body", data.Request),
	)
}

// LogResponse
func (l *Logger) LogResponse(rc string, response interface{}) {
	l.Info("Outgoing Response",
		zap.String("code", rc),
		zap.Any("body", response),
	)
}

func (l *Logger) LogDebug(message string, fields ...zap.Field) {
	l.Debug(message, fields...)
}
