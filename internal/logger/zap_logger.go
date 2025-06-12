package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

// concrete implementation using your chosen library
type zapLogger struct {
	logger *zap.Logger
	config LogConfig
}

// toZapFields converts our Field type to zap.Field and adds context fields
func (zl *zapLogger) toZapFields(ctx context.Context, fields []Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))

	// Add context fields if available
	if ctx != nil {
		if requestID, ok := ctx.Value("requestID").(string); ok {
			zapFields = append(zapFields, zap.String("requestID", requestID))
		}
	}

	// Convert custom fields to zap fields
	for _, field := range fields {
		zapFields = append(zapFields, zap.Any(field.Key, field.Value))
	}

	return zapFields
}

func (zl *zapLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	zl.logger.Debug(msg, zl.toZapFields(ctx, fields)...)
}

func (zl *zapLogger) Info(ctx context.Context, msg string, fields ...Field) {
	zl.logger.Info(msg, zl.toZapFields(ctx, fields)...)
}

func (zl *zapLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	zl.logger.Warn(msg, zl.toZapFields(ctx, fields)...)
}

func (zl *zapLogger) Error(ctx context.Context, msg string, fields ...Field) {
	zl.logger.Error(msg, zl.toZapFields(ctx, fields)...)
}

func (zl *zapLogger) Fatal(ctx context.Context, msg string, fields ...Field) {
	zl.logger.Fatal(msg, zl.toZapFields(ctx, fields)...)
}

func (zl *zapLogger) With(fields ...Field) Logger {
	return &zapLogger{
		logger: zl.logger.With(zl.toZapFields(context.Background(), fields)...),
		config: zl.config,
	}
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// New creates a new logger instance using Zap
func New(config LogConfig) (Logger, error) {
	// Set up the encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime:    CustomTimeEncoder,
	}

	if config.Development {
		// Development-friendly settings
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		encoderConfig.ConsoleSeparator = " | "
	} else {
		// Production settings
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoderConfig.EncodeCaller = nil // Disable caller in production
	}

	var encoder zapcore.Encoder
	if config.JSONFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Set up the output
	var output zapcore.WriteSyncer
	switch strings.ToLower(config.OutputPath) {
	case StdOut:
		output = zapcore.AddSync(os.Stdout)
	case StdErr:
		output = zapcore.AddSync(os.Stderr)
	default:
		file, err := os.OpenFile(config.OutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		output = zapcore.AddSync(file)
	}

	// Convert our log level to Zap's level
	var zapLevel zapcore.Level
	switch config.Level {
	case Debug:
		zapLevel = zapcore.DebugLevel
	case Info:
		zapLevel = zapcore.InfoLevel
	case Warn:
		zapLevel = zapcore.WarnLevel
	case Error:
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	core := zapcore.NewCore(encoder, output, zapLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2), zap.AddStacktrace(zapcore.ErrorLevel))

	return &zapLogger{
		logger: logger,
		config: config,
	}, nil
}
