package logger

import "context"

type Logger interface {
	Debug(ctx context.Context, msg string, fields ...Field)
	Info(ctx context.Context, msg string, fields ...Field)
	Warn(ctx context.Context, msg string, fields ...Field)
	Error(ctx context.Context, msg string, fields ...Field)
	Fatal(ctx context.Context, msg string, fields ...Field)
	With(fields ...Field) Logger
}

type Field struct {
	Key   string
	Value interface{}
}

// LogConfig holds logger configuration
type LogConfig struct {
	Level       Level
	OutputPath  string // file path or "stdout"/"stderr"
	JSONFormat  bool
	Development bool
}

// Level represents log level
type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error

	StdOut = "stdout"
	StdErr = "stderr"
)
