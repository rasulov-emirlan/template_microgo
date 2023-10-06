package postgresql

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rasulov-emirlan/template_microgo/internal/storage/postgresql/migrations"
	"github.com/rasulov-emirlan/template_microgo/internal/transport/http"
	"github.com/rasulov-emirlan/template_microgo/pkg/logging"
)

type RepoCombiner struct {
	conn *pgxpool.Pool
}

func NewRepoCombiner(ctx context.Context, dsn string, withMigrations bool, logger *slog.Logger) (RepoCombiner, error) {
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return RepoCombiner{}, fmt.Errorf("pgxpool.New: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		return RepoCombiner{}, fmt.Errorf("conn.Ping: %w", err)
	}

	if withMigrations {
		if err := migrations.Up(ctx, dsn, logging.NewGoosedLogger(logger)); err != nil {
			return RepoCombiner{}, fmt.Errorf("migrations.Up: %w", err)
		}
	}

	return RepoCombiner{conn: conn}, nil
}

func (r RepoCombiner) Check(ctx context.Context) http.HealthCheckResponse {
	res := http.HealthCheckResponse{
		Name:       "postgres",
		Status:     http.HealthStatusUp,
		IsCritical: true,
	}

	if err := r.conn.Ping(ctx); err != nil {
		res.Status = http.HealthStatusDown
		res.Message = err.Error()
	}

	return res
}
