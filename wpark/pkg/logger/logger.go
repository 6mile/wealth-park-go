package logger

import (
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// I is just a simple type alias for map[string]interface.
type I = map[string]interface{}

var (
	logger      *zap.Logger
	setupLogger sync.Once
)

// Get returns a global (singleton) zap logger instance.
func Get(tag string) *zap.Logger {
	setupLogger.Do(func() {
		level := zap.NewAtomicLevelAt(zapcore.InfoLevel)

		cfg := zap.Config{
			Encoding:         "json",
			Level:            level,
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey:    "time",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}

		switch strings.ToLower(os.Getenv("LOGLEVEL")) {
		case "debug":
			level.SetLevel(zapcore.DebugLevel)
		case "error":
			level.SetLevel(zapcore.ErrorLevel)
		}

		var err error
		logger, err = cfg.Build()
		if err != nil {
			panic("could not init zap logger")
		}

		logger.Debug("created new zap logger", zap.String("level", level.String()))
	})

	return logger.With(zap.String("tag", tag))
}
