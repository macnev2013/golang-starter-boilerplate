package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/intrigues/golang-starter-boilerplate/core"
	"github.com/intrigues/golang-starter-boilerplate/internal/api/dashboard"
	"github.com/intrigues/golang-starter-boilerplate/internal/api/login"
	"github.com/intrigues/golang-starter-boilerplate/internal/api/signup"
	"github.com/intrigues/golang-starter-boilerplate/internal/api/static"
	"github.com/intrigues/golang-starter-boilerplate/internal/api/users"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
	"github.com/intrigues/golang-starter-boilerplate/internal/db"
)

type Router struct {
	cfg    *config.Config
	logger core.Clog
	render core.Render
	db     *db.DB
}

func New(cfg *config.Config, logger core.Clog, render core.Render, db *db.DB) Router {
	return Router{
		cfg:    cfg,
		logger: logger,
		render: render,
		db:     db,
	}
}

func (router *Router) NewRouter() http.Handler {
	router.logger.Debugf("injecting routes..")

	mux := chi.NewRouter()

	// TODO: go though middleware
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(router.NoSurf)
	mux.Use(router.SessionLoad)

	mux.Get("/", static.GetHome())

	mux.Route("/login", func(r chi.Router) {
		r.Get("/", login.HandleGet(router.render))
		r.Post("/", login.HandlePost(router.render, router.logger, router.cfg, router.db))
	})
	mux.Get("/logout", login.GetLogout(router.logger, router.cfg))

	mux.Route("/signup", func(r chi.Router) {
		r.Get("/", signup.HandleGet(router.render))
		r.Post("/", signup.HandlePost(router.render, router.logger, router.cfg, router.db))
	})

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(router.Auth)
		mux.Get("/dashboard", dashboard.HandleGet(router.render))
		mux.Route("/users", func(r chi.Router) {
			r.Get("/all", users.HandleList(router.render, router.logger, router.db))
			r.Route("/{username}", func(r chi.Router) {
				r.Get("/activate", users.HandleActivate(router.render, router.logger, router.cfg, router.db))
				r.Get("/deactivate", users.HandleDeactivate(router.render, router.logger, router.cfg, router.db))
			})
		})
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
