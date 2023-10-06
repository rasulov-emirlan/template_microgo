package app

import "github.com/rasulov-emirlan/template_microgo/internal/storage/postgresql"

func (a *application) initDB() error {
	var err error

	a.pgdb, err = postgresql.NewRepoCombiner(a.ctx, a.cfg.Database.URL(), a.cfg.Flags.WithMigrations, a.logger)
	if err != nil {
		return err
	}

	a.httpChecks = append(a.httpChecks, a.pgdb.Check)

	a.logger.Info("databases initialized")

	return nil
}
