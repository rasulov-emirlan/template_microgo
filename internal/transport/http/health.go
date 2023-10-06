package http

import (
	"context"
	"net/http"
	"runtime"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

const (
	HealthStatusUp   healthStatus = "UP"
	HealthStatusDown healthStatus = "DOWN"
)

var (
	metricsAll             uint64
	metricsErrors          uint64
	metricsContextCanceled uint64
)

type (
	healthStatus string

	HealthResponse struct {
		Status  healthStatus          `json:"status"`
		Metrics HealthMetricsResponse `json:"metrics"`
		Memory  runtime.MemStats      `json:"memory"`
		Checks  []HealthCheckResponse `json:"checks"`
	}

	HealthMetricsResponse struct {
		All             uint64 `json:"all"`
		Errors          uint64 `json:"errors"`
		ContextCanceled uint64 `json:"context_canceled"`
	}

	HealthCheckResponse struct {
		Name       string       `json:"name"`
		IsCritical bool         `json:"is_critical"`
		Status     healthStatus `json:"status"`
		Message    string       `json:"message,omitempty"`
	}

	HealthCheck func(ctx context.Context) HealthCheckResponse
)

func newHealthHandler(checks []HealthCheck) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var (
			metrics = HealthMetricsResponse{
				All:             metricsAll,
				Errors:          metricsErrors,
				ContextCanceled: metricsContextCanceled,
			}
			checkResults = make([]HealthCheckResponse, len(checks))
		)

		status := HealthStatusUp

		for i, check := range checks {
			checkResults[i] = check(ctx)
			if checkResults[i].Status == HealthStatusDown && checkResults[i].IsCritical {
				status = HealthStatusDown
			}
		}

		res := HealthResponse{
			Status:  status,
			Metrics: metrics,
			Checks:  checkResults,
		}

		runtime.ReadMemStats(&res.Memory)

		return c.JSON(http.StatusOK, res)
	}
}

func middlewareMetrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		atomic.AddUint64(&metricsAll, 1)

		err := next(c)

		if err != nil {
			atomic.AddUint64(&metricsErrors, 1)
		}

		if c.Request().Context().Err() != nil {
			atomic.AddUint64(&metricsContextCanceled, 1)
		}

		return err
	}
}
