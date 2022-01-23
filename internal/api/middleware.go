package api

import (
	"net/http"

	"github.com/intrigues/golang-starter-boilerplate/internal/config"
	"github.com/intrigues/golang-starter-boilerplate/internal/models"
	"github.com/justinas/nosurf"
)

// NoSurf is the csrf protection middleware
func (router *Router) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		// TODO - get this from env
		HttpOnly: false,
		Path:     "/",
		// TODO - get this from env
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func (router *Router) SessionLoad(next http.Handler) http.Handler {
	return router.cfg.Session.LoadAndSave(next)
}

func (router *Router) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(r, router.cfg) {
			router.cfg.Session.Put(r.Context(), "error", "Please login to continue")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isAuthenticated(r *http.Request, cfg *config.Config) bool {
	exists := cfg.Session.Exists(r.Context(), "currentuser")
	if exists {
		currentUser := cfg.Session.Get(r.Context(), "currentuser").(models.Users)
		return !(currentUser.Status == 0)
	}
	return exists
}
