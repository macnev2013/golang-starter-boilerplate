package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/intrigues/golang-starter-boilerplate/core"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
	"github.com/intrigues/golang-starter-boilerplate/internal/db"
	"github.com/intrigues/golang-starter-boilerplate/internal/forms"
	"github.com/intrigues/golang-starter-boilerplate/internal/models"
)

func HandleList(render core.Render, logger core.Clog, db *db.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var users []models.Users
		db.Conn.Find(&users)

		data := make(map[string]interface{})
		data["users"] = users
		render.RenderTemplate(rw, r, "users.page.tmpl", &models.TemplateData{
			Form: forms.New(nil),
			Data: data,
		})
	}
}

func HandleActivate(render core.Render, logger core.Clog, cfg *config.Config, db *db.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		logger.Debugf("activating user:", username)

		var user models.Users
		db.Conn.First(&user, "username = ?", username)
		user.Status = 1
		db.Conn.Save(&user)
		cfg.Session.Put(r.Context(), "flash", "User activated")
		http.Redirect(rw, r, "/admin/users/all", http.StatusSeeOther)
	}
}

func HandleDeactivate(render core.Render, logger core.Clog, cfg *config.Config, db *db.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		logger.Debugf("deactivating user:", username)
		var user models.Users
		db.Conn.First(&user, "username = ?", username)
		user.Status = 0
		db.Conn.Save(&user)
		cfg.Session.Put(r.Context(), "flash", "User activated")
		http.Redirect(rw, r, "/admin/users/all", http.StatusSeeOther)
	}
}
