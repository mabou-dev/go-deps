package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/mabou-dev/go-deps/pkg/core/golang"
	"github.com/mabou-dev/go-deps/pkg/core/npm"
)

// Config holds the application configuration
type Config struct {
	App       AppConfig                  `mapstructure:"app"`
	Technical map[string]TechnicalConfig `mapstructure:"technical"`
}

// AppConfig holds application-specific settings
type AppConfig struct {
	Name   string     `mapstructure:"name"`
	Debug  bool       `mapstructure:"debug"`
	Logger zap.Config `mapstructure:"logger"`
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
			Logger: zap.Config{
				Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
				Development: false,
				Encoding:    "json",
				EncoderConfig: zapcore.EncoderConfig{
					TimeKey:        "timestamp",
					LevelKey:       "level",
					NameKey:        "logger",
					CallerKey:      "caller",
					MessageKey:     "msg",
					StacktraceKey:  "stacktrace",
					LineEnding:     zapcore.DefaultLineEnding,
					EncodeLevel:    zapcore.LowercaseLevelEncoder,
					EncodeTime:     zapcore.ISO8601TimeEncoder,
					EncodeDuration: zapcore.StringDurationEncoder,
					EncodeCaller:   zapcore.ShortCallerEncoder,
				},
				OutputPaths:      []string{"stdout"},
				ErrorOutputPaths: []string{"stderr"},
			},
		},
		Technical: map[string]TechnicalConfig{
			npm.NAME: {
				Enabled:     true,
				RegistryURL: "https://registry.npmjs.org/",
			},
			golang.NAME: {
				Enabled:     true,
				RegistryURL: "https://proxy.golang.org/",
			},
			//sbt.NAME: {
			//	Enabled:     true,
			//	RegistryURL: "https://repo1.maven.org/maven2/",
			//},
			//python.NAME: {
			//	Enabled:     true,
			//	RegistryURL: "https://pypi.org/simple",
			//},
			// Add other technologies with their default configs here
		},
	}
}

// LoadConfig loads the configuration from a JSON file with fallback to defaults
func LoadConfig() (*Config, error) {
	config := DefaultConfig()
	// Try to read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, return default config
			return config, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal config into struct, orverride defaults
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
