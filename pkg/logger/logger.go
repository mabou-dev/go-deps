package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger defines a minimal logging interface independent of any implementation
type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
}

// zapLoggerAdapter wraps *zap.Logger to implement our Logger interface
type zapLoggerAdapter struct {
	logger *zap.Logger
}

// NewLogger creates and returns a new Logger instance backed by Zap
func NewLogger(debug bool) Logger {
	var config zap.Config

	if debug {
		config = zap.NewDevelopmentConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	} else {
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	zapLogger, err := config.Build()
	if err != nil {
		fmt.Println("Error creating zap logger:", err)
		zapLogger = zap.NewNop() // Fallback to no-op logger
		return &zapLoggerAdapter{logger: zapLogger}
	}

	return &zapLoggerAdapter{logger: zapLogger}
}

// Info logs an info message with key-value pairs
func (z *zapLoggerAdapter) Info(msg string) {
	z.logger.Info(msg)
}

// Debug logs a debug message with key-value pairs
func (z *zapLoggerAdapter) Debug(msg string) {
	z.logger.Debug(msg)
}

// Error logs an error message with key-value pairs
func (z *zapLoggerAdapter) Error(msg string) {
	z.logger.Error(msg)
}

// SyncLogger flushes any buffered log entries
func SyncLogger(log Logger) error {
	// Cast to adapter to access underlying zap logger
	if adapter, ok := log.(*zapLoggerAdapter); ok {
		return adapter.logger.Sync()
	}
	return nil
}
