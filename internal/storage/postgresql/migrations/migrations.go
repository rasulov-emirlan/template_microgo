package migrations

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var fs embed.FS

func Up(ctx context.Context, dsn string, logger goose.Logger) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose.SetDialect: %w", err)
	}

	goose.SetBaseFS(fs)

	goose.SetLogger(logger)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	if err := goose.Up(conn, "."); err != nil {
		return fmt.Errorf("goose.Up: %w", err)
	}

	return nil
}
