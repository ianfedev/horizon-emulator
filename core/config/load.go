package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

// LoadConfig loads configuration from the given file path and unmarshalls it into the Config struct.
func LoadConfig(path string, logger *zap.Logger) (*Config, error) {

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("ini")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var tempCfg Config
	if err := v.Unmarshal(&tempCfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	CheckSecurityAlerts(&tempCfg, logger)

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	SetDefaultValues(v)

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil

}
