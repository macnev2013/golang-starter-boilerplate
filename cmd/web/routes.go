package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
	"github.com/intrigues/golang-starter-boilerplate/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.GetHome)

	mux.Get("/login", handlers.Repo.GetLogin)

	mux.Get("/dashboard", handlers.Repo.GetDashboard)

	mux.Get("/users", handlers.Repo.GetUsers)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
