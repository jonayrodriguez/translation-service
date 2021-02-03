package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logging represents the logging configuration.
type Logging struct {
	Level       string
	Development bool
}

// Logger the custom smp logger
var Logger *zap.Logger

// Config configures the logger
func Config(config Logging) (*zap.Logger, error) {
	if Logger == nil {
		var level zapcore.Level
		err := level.UnmarshalText([]byte(config.Level))
		if err != nil {
			level = zapcore.InfoLevel
		}
		core := zapcore.NewCore(
			getEncoder(),
			zapcore.AddSync(os.Stdout),
			level,
		)
		Logger = zap.New(core, zap.AddCaller())
		Logger = Logger.With(zap.Namespace("fields"))

		if config.Development {
			_ = zap.Development()
		}
		zap.ReplaceGlobals(Logger)
	}
	return Logger, nil
}

// Get the encoder in a specific format
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		MessageKey:     "msg",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}
