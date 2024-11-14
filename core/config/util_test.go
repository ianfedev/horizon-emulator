package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// createTestLogger creates a logger that writes logs to a buffer for verification.
func createTestLogger() (*zap.Logger, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(buf),
		zap.InfoLevel,
	)
	logger := zap.New(core)
	return logger, buf
}

// verifySecurityLog checks if the expected security alert is logged for specific fields.
func verifySecurityLog(t *testing.T, logOutput string, fields []string) {
	t.Helper()
	assert.Contains(t, logOutput, "Security Alert: Sensitive data detected in configuration")
	for _, field := range fields {
		assert.Contains(t, logOutput, field)
	}
}

// TestSetDefaultValues performs a quick test to check if config is being assigned
// correctly via reflection.
func TestSetDefaultValues(t *testing.T) {
	v := viper.New()
	cfg := Config{}

	SetDefaultValues(v, "config", cfg)
	assert.Equal(t, "127.0.0.1", v.GetString("config.server.ip"), "Default value for server.ip should be 127.0.0.1")
	assert.Equal(t, "example_db", v.GetString("config.database.name"), "Default value for database.name should be example_db")
	assert.Equal(t, true, v.GetBool("config.logging.console_color"), "Default value for logging.console_color should be true")

}

// TestCheckSecurityAlerts test if the environment variable message is logged at console.
func TestCheckSecurityAlerts(t *testing.T) {
	logger, buf := createTestLogger()

	config := &Config{
		Server: ServerConfig{
			IP:          "127.0.0.1",
			Port:        8080,
			Environment: "PRODUCTION",
		},
		Database: DatabaseConfig{
			Name:     "prod_db",
			Password: "prod_secret",
			Username: "admin",
			Host:     "db.prod.example.com",
			Port:     5432,
		},
	}

	CheckSecurityAlerts(config, logger)
	verifySecurityLog(t, buf.String(), []string{"Password", "Username", "Host"})
}

// TestCheckStruct test if the environment variable message is logged at console.
func TestCheckStruct(t *testing.T) {
	logger, buf := createTestLogger()

	config := &Config{
		Server: ServerConfig{
			IP:          "127.0.0.1",
			Port:        8080,
			Environment: "PRODUCTION",
		},
		Database: DatabaseConfig{
			Name:     "prod_db",
			Password: "prod_secret",
			Username: "admin",
			Host:     "db.prod.example.com",
			Port:     5432,
		},
	}

	// Run the checkStruct function directly
	checkStruct(reflect.ValueOf(config), "PRODUCTION", logger)
	verifySecurityLog(t, buf.String(), []string{"Password", "Username", "Host"})
}

// TestCheckStructNonStruct verifies that checkStruct correctly handles non-struct values.
func TestCheckStructNonStruct(t *testing.T) {
	logger, buf := createTestLogger()

	// Pass a non-struct value, such as an integer
	nonStructValue := 42
	checkStruct(reflect.ValueOf(nonStructValue), "PRODUCTION", logger)

	// Ensure that no log entries are created since it's not a struct
	if buf.String() != "" {
		t.Errorf("Expected no log output for non-struct value, got: %s", buf.String())
	}
}

// TestCreateDefaultConfig checks if configuration file is created when missing.
func TestCreateDefaultConfig(t *testing.T) {

	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.ini")

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err := CreateDefaultConfig(configPath, logger)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatalf("expected config file to be created, but it does not exist")
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("ini")

	if err := v.ReadInConfig(); err != nil {
		t.Fatalf("expected to read config file, but got error: %v", err)
	}

	fmt.Println("TestCreateDefaultConfig passed")
}

func TestCreateDefaultConfig_Error(t *testing.T) {
	logger, _ := createTestLogger()

	invalidPath := "/invalid/path/config.ini"

	err := CreateDefaultConfig(invalidPath, logger)

	if err == nil {
		t.Fatalf("Expected an error when creating default config file, but got nil")
	}

	expectedErrorMessage := "error creating default config file"
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("Expected error message to contain %q, but got: %v", expectedErrorMessage, err)
	}

}
