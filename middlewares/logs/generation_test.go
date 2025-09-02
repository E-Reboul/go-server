package logs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateLogDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "logs")

	err := createLogDirectory()
	if err != nil {
		t.Fatalf("Failed to create log directory: %v", err)
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Errorf("Log directory was not created")
	}
}
