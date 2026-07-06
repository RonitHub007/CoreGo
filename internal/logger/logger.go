package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Options holds configuration for our logger.
type Options struct {
	Level zapcore.Level
}

// Option is a functional option type for configuring the logger.
// This is a classic functional programming pattern in Go.
type Option func(*Options)

// WithLevel is a functional option that sets the logging level.
func WithLevel(level zapcore.Level) Option {
	return func(o *Options) {
		o.Level = level
	}
}

// New creates a new structured zap logger using the provided options.
func New(opts ...Option) (*zap.Logger, error) {
	// 1. Set default options
	options := &Options{
		Level: zapcore.InfoLevel,
	}

	// 2. Apply any passed-in functional options
	for _, opt := range opts {
		opt(options)
	}

	// 3. Build the zap logger
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(options.Level)

	// You can customize output formats here (e.g., JSON vs Console).
	// ProductionConfig defaults to JSON, which is perfect for parsing by log aggregators.
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
