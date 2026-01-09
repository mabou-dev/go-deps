package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	App AppConfig `mapstructure:"app"`
}

// AppConfig holds application-specific settings
type AppConfig struct {
	Name  string `mapstructure:"name"`
	Debug bool   `mapstructure:"debug"`
}

// DefaultConfig returns a default configuration structure
func DefaultConfig() *Config {
	return &Config{
		App: AppConfig{
			Name:  "go-deps",
			Debug: false,
		},
	}
}

// LoadConfig loads the configuration from a JSON file with fallback to defaults
func LoadConfig() (*Config, error) {
	// Try to read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, return default config
			return DefaultConfig(), nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
