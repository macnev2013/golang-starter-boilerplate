package server

import (
	"net/http"

	"github.com/intrigues/golang-starter-boilerplate/core"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
)

func ListenAndServe(cfg config.Config, router http.Handler, logger core.Clog) error {
	// TODO: retrive binding from the config
	// TODO: add this in go routine
	// TODO: add option to run with https
	srv := &http.Server{
		Addr:    "localhost:" + cfg.Port,
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
