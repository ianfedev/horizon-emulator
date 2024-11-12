package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"reflect"
)

// SetDefaultValues sets default values for the configuration.
func SetDefaultValues(v *viper.Viper) {
	v.SetDefault("server.ip", "127.0.0.1")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.environment", "development")

	v.SetDefault("database.name", "horizon")
	v.SetDefault("database.username", "user")
	v.SetDefault("database.password", "password")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 3306)

	v.SetDefault("logging.console_color", false)
	v.SetDefault("logging.json", false)
	v.SetDefault("logging.level", "info")
}

// CheckSecurityAlerts recurses through the configuration structure to detect fields marked with 'security_alert' tags
// in the specified environment. If such a field is found, it logs a security alert using the zap logger.
func CheckSecurityAlerts(c *Config, logger *zap.Logger) {
	checkStruct(reflect.ValueOf(c), c.Server.Environment, logger)
}

func checkStruct(v reflect.Value, env string, logger *zap.Logger) {

	// Handle pointers by getting the actual value they point to
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Only proceed if the value is a struct
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		securityAlertTag := fieldType.Tag.Get("security_alert")
		if securityAlertTag != "" && securityAlertTag == string(env) {
			// If the field is set (non-zero value), log a security alert
			if !field.IsZero() {
				logger.Info("Security Alert: Sensitive data detected in configuration",
					zap.String("field", fieldType.Name),
					zap.String("environment", string(env)))
			}
		}

		// Recursively check nested structs
		if field.Kind() == reflect.Struct || field.Kind() == reflect.Ptr {
			checkStruct(field, env, logger)
		}
	}

}

// CreateDefaultConfig creates a default config file if it does not exist.
func CreateDefaultConfig(path string, logger *zap.Logger) error {

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("ini")

	SetDefaultValues(v)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := v.WriteConfigAs(path); err != nil {
			return fmt.Errorf("error creating default config file: %w", err)
		}
		logger.Info("Config file not found. Created default config file.")
	}

	return nil
}
