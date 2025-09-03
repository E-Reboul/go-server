package logs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateLogDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "logs")

	err := CreateLogDirectory(logPath)
	if err != nil {
		t.Fatalf("Failed to create log directory: %v", err)
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Errorf("Log directory was not created.")
	}
}

func TestCreateLoggerDirectories(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "logs")
	loggersPaths := getLogsCategoriesPaths()

	err := CreateLogDirectory(logPath)
	if err != nil {
		t.Fatalf("Failed to create log directory: %v", err)
	}

	err = CreateLoggersDirectories(loggersPaths)
	if err != nil {
		t.Fatalf("Failed to create loggers directories: %v", err)
	}

	files, err := os.ReadDir(logPath)
	if err != nil {
		t.Fatalf("Failed to read log directory: %v", err)
	}
	for _, file := range files {
		t.Logf("Log file created: %s", file.Name())
	}
}

func TestCreateLoggerFile(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "logs")
	loggersPaths := getLogsCategoriesPaths()

	err := CreateLogDirectory(logPath)
	if err != nil {
		t.Fatalf("Failed to create log directory: %v", err)
	}

	err = CreateLoggersDirectories(loggersPaths)
	if err != nil {
		t.Fatalf("Failed to create loggers directories: %v", err)
	}

	for category, path := range loggersPaths {
		_ = os.Remove(path)
		_ = CreateLoggerFile(category)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Logger file for category %s was not created.", category)
		}
	}
}
