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
	err := run()
	if err != nil {
		fmt.Printf("Error in starting application %v", err)
	}
}

func run() error {
	// initiating configuratio
	cfg := config.Load()

	// inititating loggers
	logger := clog.NewLogger(cfg)

	logger.Debugf("creating new session")
	session := session.New()
	cfg.Session = session
	logger.Debugf("session initialized", cfg.Session)

	// initiating renderer
	logger.Debugf("initiating render")
	render := render.New(&cfg, logger)
	
	// initiating database connections
	logger.Debugf("initiating database connections")
	// initializing database
	db, err := db.NewDBConnection(cfg, logger)
	if err != nil {
		logger.Errorf("error in initializing database: %v", err)
	}

	api := api.New(&cfg, logger, render, db)
	router := api.NewRouter()

	// add this non blocking call
	if err := server.ListenAndServe(cfg, router, logger); err != nil {
		logger.Errorf("error starting http server: %v", err)
	}
	return nil
}
