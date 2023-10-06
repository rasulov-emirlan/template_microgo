package logging

import (
	"log/slog"
	"os"

	"github.com/pressly/goose/v3"
)

type GoosedLogger struct {
	logger *slog.Logger
}

var _ goose.Logger = (*GoosedLogger)(nil)

func NewGoosedLogger(logger *slog.Logger) *GoosedLogger {
	return &GoosedLogger{logger: logger}
}

func (l *GoosedLogger) Fatalf(format string, v ...any) {
	attrs := make([]any, len(v))
	for i, val := range v {
		attrs[i] = slog.Any("v", val)
	}
	l.logger.Error(format, attrs...)
	os.Exit(1)
}

func (l *GoosedLogger) Printf(format string, v ...any) {
	attrs := make([]any, len(v))
	for i, val := range v {
		attrs[i] = slog.Any("v", val)
	}
	l.logger.Info(format, attrs...)
}
