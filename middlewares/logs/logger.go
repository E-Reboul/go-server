package logs

import (
	"go.uber.org/zap"
)

// Category type
type LogCategory string

// Logger at category
var Loggers = map[LogCategory]*zap.Logger{}

func WriteInfo(category LogCategory, msg string, fields ...zap.Field) {
	logger := Loggers[category]
	if logger != nil {
		logger.Info(msg, fields...)
	}
}

func WriteError(category LogCategory, msg string, fields ...zap.Field) {
	logger := Loggers[category]
	if logger != nil {
		logger.Error(msg, fields...)
	}
}

func WriteWarn(category LogCategory, msg string, fields ...zap.Field) {
	logger := Loggers[category]
	if logger != nil {
		logger.Warn(msg, fields...)
	}
}

func WriteDebug(category LogCategory, msg string, fields ...zap.Field) {
	logger := Loggers[category]
	if logger != nil {
		logger.Debug(msg, fields...)
	}
}
