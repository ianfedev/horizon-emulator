package config

import (
	"bytes"
	"reflect"
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

func TestSetDefaultValues(t *testing.T) {
	v := viper.New()
	SetDefaultValues(v)

	assert.Equal(t, "127.0.0.1", v.GetString("server.ip"))
	assert.Equal(t, 8080, v.GetInt("server.port"))
	assert.Equal(t, "development", v.GetString("server.environment"))
	assert.Equal(t, "horizon", v.GetString("database.name"))
	assert.Equal(t, "user", v.GetString("database.username"))
	assert.Equal(t, "password", v.GetString("database.password"))
	assert.Equal(t, "localhost", v.GetString("database.host"))
	assert.Equal(t, 3306, v.GetInt("database.port"))
	assert.False(t, v.GetBool("logging.console_color"))
	assert.False(t, v.GetBool("logging.json"))
	assert.Equal(t, "info", v.GetString("logging.level"))
}

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
