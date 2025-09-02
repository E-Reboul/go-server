package logs

/*
 * Log writing functions for the API server.
 * This file contains WriteInfo, WriteError, etc. to write logs by category.
 */

import (
	types "github.com/E-Reboul/go-server/types/logs"
	"go.uber.org/zap"
)

func WriteInfo(category types.LogCategory, msg string, fields ...zap.Field) {
	logger := Loggers[category]
	if logger != nil {
		logger.Info(msg, fields...)
	}
}

func WriteError(category types.LogCategory, msg string, fields ...zap.Field) {
	logger := Loggers[category]
	if logger != nil {
		logger.Error(msg, fields...)
	}
}

func WriteWarn(category types.LogCategory, msg string, fields ...zap.Field) {
	logger := Loggers[category]
	if logger != nil {
		logger.Warn(msg, fields...)
	}
}

func WriteDebug(category types.LogCategory, msg string, fields ...zap.Field) {
	logger := Loggers[category]
	if logger != nil {
		logger.Debug(msg, fields...)
	}
}
