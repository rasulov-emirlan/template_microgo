package http

import (
	"context"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rasulov-emirlan/template_microgo/config"
)

type server struct {
	srvr    *http.Server
	network string
}

func NewServer(cfg config.Config, checks []HealthCheck) (server, error) {
	router := echo.New()

	router.Use(middleware.Gzip())
	router.Use(middleware.RemoveTrailingSlash())
	router.Use(middlewareMetrics)

	router.GET("/health", newHealthHandler(checks))

	srvr := http.Server{
		Addr:         cfg.HttpAddr,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		Handler:      router,
	}

	return server{srvr: &srvr, network: cfg.HttpNetwork}, nil
}

func (s server) Start() error {
	listener, err := net.Listen(s.network, s.srvr.Addr)
	if err != nil {
		return err
	}

	if err := s.srvr.Serve(listener); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s server) Stop(ctx context.Context) error {
	return s.srvr.Shutdown(ctx)
}
