package config

import (
	"go.uber.org/zap"
	"path/filepath"
	"testing"
)

// TestLoadConfig creates a demo configuration expecting to be load without problem.
func TestLoadConfig(t *testing.T) {

	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.ini")

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err := CreateDefaultConfig(configPath, logger)
	if err != nil {
		t.Fatalf("expected no error creating default config, but got %v", err)
	}

	loadedConfig, err := LoadConfig(configPath, logger)
	if err != nil {
		t.Fatalf("expected no error loading config, but got %v", err)
	}

	if loadedConfig == nil {
		t.Fatalf("expected config to be loaded, but got nil")
	}

}
