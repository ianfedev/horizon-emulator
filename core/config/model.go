package config

// Environment defines the type for server environment configurations.
// It can be one of the following: PRODUCTION, DEVELOPMENT, or STAGING.
type Environment string

const (
	PRODUCTION  Environment = "production"  // PRODUCTION represents a live/production environment.
	DEVELOPMENT Environment = "development" // DEVELOPMENT represents a development environment.
	STAGING     Environment = "staging"     // STAGING represents a staging/testing environment.
)

// ServerConfig holds the configuration for the server settings.
// It includes IP address, port number, and environment type.
type ServerConfig struct {
	IP          string      `mapstructure:"ip"`          // IP address of the server.
	Port        int         `mapstructure:"port"`        // Port number on which the server listens.
	Environment Environment `mapstructure:"environment"` // Environment type (PRODUCTION, DEVELOPMENT, STAGING).
}

// DatabaseConfig holds the configuration for the database settings.
// It includes the database name, password, username, host, and port.
type DatabaseConfig struct {
	Name     string `mapstructure:"db_name" security_alert:"PRODUCTION"`     // Name of the database.
	Password string `mapstructure:"db_password" security_alert:"PRODUCTION"` // Password for the database user.
	Username string `mapstructure:"db_username" security_alert:"PRODUCTION"` // Username for the database.
	Host     string `mapstructure:"db_host" security_alert:"PRODUCTION"`     // Host address of the database.
	Port     int    `mapstructure:"db_port" security_alert:"PRODUCTION"`     // Port number on which the database listens.
}

// LoggingConfig holds the configuration for the logging system.
type LoggingConfig struct {
	ConsoleColor bool   `mapstructure:"console_color"` // ConsoleColor if console should be colored or not.
	JSON         bool   `mapstructure:"json"`          // JSON if file logging should be in json for monitoring.
	Level        string `mapstructure:"level"`         // Level of the logging.
}

// Config is the root configuration structure that holds both server and database configurations.
// It contains the configuration details for initializing the server and connecting to the database.
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`   // Server configuration.
	Database DatabaseConfig `mapstructure:"database"` // Database configuration.
	Logging  LoggingConfig  `mapstructure:"logging"`  // Logging configuration.
}

// NOTE: Remember to add the default value to the utility function
