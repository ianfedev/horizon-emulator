package config

// ServerConfig holds the configuration for the server settings.
// It includes IP address, port number, and environment type.
type ServerConfig struct {
	IP          string `mapstructure:"ip" default:"127.0.0.1"`            // IP address of the server.
	Port        int    `mapstructure:"port" default:"8080"`               // Port number on which the server listens.
	Environment string `mapstructure:"environment" default:"DEVELOPMENT"` // Environment type (PRODUCTION, DEVELOPMENT, STAGING).
}

// DatabaseConfig holds the configuration for the database settings.
// It includes the database name, password, username, host, and port.
type DatabaseConfig struct {
	Name     string `mapstructure:"name" security_alert:"PRODUCTION" default:"example_db"`   // Name of the database.
	Password string `mapstructure:"password" security_alert:"PRODUCTION" default:"password"` // Password for the database user.
	Username string `mapstructure:"username" security_alert:"PRODUCTION" default:"user"`     // Username for the database.
	Host     string `mapstructure:"host" security_alert:"PRODUCTION" default:"localhost"`    // Host address of the database.
	Port     int    `mapstructure:"port" security_alert:"PRODUCTION" default:"5432"`         // Port number on which the database listens.
}

// LoggingConfig holds the configuration for the logging system.
type LoggingConfig struct {
	ConsoleColor bool   `mapstructure:"console_color" default:"true"` // ConsoleColor if console should be colored or not.
	JSON         bool   `mapstructure:"json" default:"false"`         // JSON if file logging should be in json for monitoring.
	Level        string `mapstructure:"level" default:"INFO"`         // Level of the logging.
}

// Config is the root configuration structure that holds both server and database configurations.
// It contains the configuration details for initializing the server and connecting to the database.
type Config struct {
	Server   ServerConfig   `mapstructure:"server" default:""`   // Server configuration.
	Database DatabaseConfig `mapstructure:"database" default:""` // Database configuration.
	Logging  LoggingConfig  `mapstructure:"logging" default:""`  // Logging configuration.
}

// NOTE: Remember to add the default value to the utility function.
// Also, remember to add corresponding tests.
