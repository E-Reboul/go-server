package logs

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ServerCategory      LogCategory = "server"
	HealthcheckCategory LogCategory = "healthcheck"
)

// All logs folders
var logsDirectoriesPaths = map[LogCategory]string{
	HealthcheckCategory: filepath.Join("logs", string(HealthcheckCategory), string(ServerCategory)+".log"),
	ServerCategory:      filepath.Join("logs", string(ServerCategory), string(ServerCategory)+".log"),
}

func createLogDirectory() {
	// Get logs directory informations, if not exist create it else do nothing
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

func CreateLoggerFile(category LogCategory) zapcore.Core {
	// get all path logs folder by categories
	path := logsDirectoriesPaths[category]
	// Initialize zap encoder
	cfg := zap.NewProductionEncoderConfig()
	// Set format time fields to timestamp
	cfg.TimeKey = "timestamp"
	// Formated timestamp on ISO8601TimeEncoder
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	// Used JSON
	encoder := zapcore.NewConsoleEncoder(cfg)
	// Open file in write/create/read permissions
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// If error on write/create log file, write to console
	if err != nil {
		return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	}
	// Else write on
	return zapcore.NewCore(encoder, zapcore.AddSync(file), zap.DebugLevel)
}

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
