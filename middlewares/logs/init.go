package logs

/*
 * Logger initialization for the API server.
 * This file contains functions to initialize and close loggers.
 */

import (
	"go.uber.org/zap"
)

func LoadLogger() {
	createLogDirectory()
	createLoggerDirectories()
	for category := range logsDirectoriesPaths {
		core := CreateLoggerFile(category)
		logger := zap.New(core)
		Loggers[category] = logger
	}
}

// Ferme tous les loggers (flush)
func CloseLoggers() {
	for _, logger := range Loggers {
		_ = logger.Sync()
	}
}
