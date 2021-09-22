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
	mux.Post("/login", handlers.Repo.PostLogin)

	mux.Get("/signup", handlers.Repo.GetSignup)
	mux.Post("/signup", handlers.Repo.PostSignup)
	mux.Get("/logout", handlers.Repo.GetLogout)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", handlers.Repo.GetDashboard)
		mux.Route("/users", func(r chi.Router) {
			r.Get("/all", handlers.Repo.GetUsers)
			r.Route("/{username}", func(r chi.Router) {
				r.Get("/activate", handlers.Repo.ActivateUser)
				r.Get("/deactivate", handlers.Repo.DeactivateUser)
			})
		})
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
