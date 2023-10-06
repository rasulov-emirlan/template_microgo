package app

import "github.com/rasulov-emirlan/template_microgo/internal/transport/http"

func (a *application) initHTTP() error {
	srvr, err := http.NewServer(a.cfg, a.httpChecks)
	if err != nil {
		return err
	}

	a.closings = append(a.closings, func() {
		if err := srvr.Stop(a.ctx); err != nil {
			a.logger.Error("http server stop", err)
		}
	})

	go func() {
		if err := srvr.Start(); err != nil {
			a.fatal("http server", err)
		}
	}()

	a.logger.Info("http server initialized", "addr", a.cfg.HttpAddr)

	return nil
}
