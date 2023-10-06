package logging

import (
	"context"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

type ReqKey string

const (
	ReqID ReqKey = "req_id"
)

type contextHandler struct {
	slog.Handler
}

func (h contextHandler) Handle(ctx context.Context, r slog.Record) error {
	reqID := ctx.Value(ReqID)
	if reqID != nil {
		reqIdString, ok := reqID.(string)
		if ok {
			panic("req_id is not a string")
		}
		r.AddAttrs(slog.String("req_id", reqIdString))
	}
	return h.Handler.Handle(ctx, r)
}

var levels = map[string]slog.Level{
	"dev":   slog.LevelDebug,
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func NewLogger(logLevel string) *slog.Logger {
	var l *slog.Logger
	var h slog.Handler
	lvl, found := levels[logLevel]
	if !found {
		lvl = slog.LevelDebug
	}
	switch logLevel {
	case "dev":
		handler := tint.NewHandler(os.Stdout, &tint.Options{
			Level: lvl,
		})
		h = contextHandler{handler}
	default:
		h = contextHandler{slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: lvl,
		})}
	}

	l = slog.New(h)

	slog.SetDefault(l)
	return l
}
