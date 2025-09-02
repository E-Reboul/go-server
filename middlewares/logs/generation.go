package logs

/*
 * Log directory and file generation utilities for the API server.
 * This file contains helper functions to create log directories and files.
 */

import (
	"os"
	"path/filepath"

	types "github.com/E-Reboul/go-server/types/logs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLogDirectory() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		mkdirErr := os.Mkdir("logs", 0755)
		if mkdirErr != nil {
			println("Failed to create directory: ", err.Error())
		}
	}
}

func createLoggerDirectories() {
	// For all loggers paths create associated directory
	for _, path := range logsDirectoriesPaths {
		dir := filepath.Dir(path)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			println("Failed dto create log directory:", dir, err.Error())
		}
	}
}

func CreateLoggerFile(category types.LogCategory) zapcore.Core {
	// Get the log file path for the given category
	path := logsDirectoriesPaths[category]
	// Initialize zap encoder configuration
	cfg := zap.NewProductionEncoderConfig()
	// Set the time field format to timestamp
	cfg.TimeKey = "timestamp"
	// Format timestamps using ISO8601
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	// Use console encoder (can be changed to JSON if needed)
	encoder := zapcore.NewConsoleEncoder(cfg)
	// Open the log file with append, create, and write permissions
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// If the log file cannot be opened, write logs to the console
	if err != nil {
		return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	}
	// Otherwise, write logs to the file
	return zapcore.NewCore(encoder, zapcore.AddSync(file), zap.DebugLevel)
}
