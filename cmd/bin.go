package cmd

import (
	"fmt"

	"github.com/intrigues/golang-starter-boilerplate/internal/api"
	"github.com/intrigues/golang-starter-boilerplate/internal/clog"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
	"github.com/intrigues/golang-starter-boilerplate/internal/db"
	"github.com/intrigues/golang-starter-boilerplate/internal/render"
	"github.com/intrigues/golang-starter-boilerplate/internal/server"
	"github.com/intrigues/golang-starter-boilerplate/internal/session"
)

func Init() {
	err := run
	if err != nil {
		fmt.Println("Error in starting application")
	}
}

func run() error {
	// initiating configuration
	cfg := config.Load()

	// inititating loggers
	logger := clog.NewLogger(cfg)

	// initiating database connections
	logger.Debugf("initiating database connections")

	// initiating renderer
	render := render.New(cfg, logger)

	// initializing database
	db, err := db.NewDBConnection(cfg, logger)
	if err != nil {
		logger.Errorf("error in initializing database: %v", err)
	}

	session := session.New()
	cfg.Session = session

	api := api.New(&cfg, logger, render, db)
	router := api.NewRouter()
	// add this non blocking call
	if err := server.ListenAndServe(cfg, router, logger); err != nil {
		logger.Errorf("error starting http server: %v", err)
	}
	return nil
}
