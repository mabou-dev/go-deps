package config

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/mabou-dev/go-deps/pkg/technology/npm"
)

// Config holds the application configuration
type Config struct {
	App       AppConfig                  `mapstructure:"app"`
	Technical map[string]TechnicalConfig `mapstructure:"technical"`
}

// AppConfig holds application-specific settings
type AppConfig struct {
	Name  string `mapstructure:"name"`
	Debug bool   `mapstructure:"debug"`
}

type TechnicalConfig struct {
	Enabled     bool   `mapstructure:"enabled"`
	RegistryURL string `mapstructure:"registry_url"`
}

// DefaultConfig returns a default configuration structure
func DefaultConfig() *Config {
	return &Config{
		App: AppConfig{
			Name:  "go-deps",
			Debug: false,
		},
		Technical: map[string]TechnicalConfig{
			npm.NAME: {
				Enabled:     true,
				RegistryURL: "https://registry.npmjs.org/",
			},
			// Add other technologies with their default configs here
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
