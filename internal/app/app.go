package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/rasulov-emirlan/template_microgo/config"
	"github.com/rasulov-emirlan/template_microgo/internal/storage/postgresql"
	"github.com/rasulov-emirlan/template_microgo/internal/transport/http"
	"github.com/rasulov-emirlan/template_microgo/pkg/logging"
)

type application struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	closings  []func()

	cfg config.Config

	logger *slog.Logger

	pgdb postgresql.RepoCombiner

	httpChecks []http.HealthCheck
}

func Run() {
	a := application{}

	a.ctx, a.ctxCancel = context.WithCancel(context.Background())
	defer a.ctxCancel()

	a.closings = append(a.closings, a.ctxCancel)

	cfg, err := config.LoadConfigs()
	if err != nil {
		panic(err)
	}

	a.cfg = cfg

	a.logger = logging.NewLogger(cfg.LogLevel)

	if err := a.initDB(); err != nil {
		a.fatal("initDB", err)
	}

	if err := a.initHTTP(); err != nil {
		a.fatal("initHTTP", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println() // new line after ^C

	a.cleanup()
}

func (a application) fatal(msg string, err error) {
	a.logger.Error(msg, slog.String("err", err.Error()))
	a.cleanup()
	os.Exit(-1)
}

func (a application) cleanup() {
	a.logger.Info("cleanup")
	for _, f := range a.closings {
		f()
	}
	a.logger.Info("cleanup done")
}
