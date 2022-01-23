package login

import (
	"net/http"

	"github.com/intrigues/golang-starter-boilerplate/core"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
	"github.com/intrigues/golang-starter-boilerplate/internal/db"
	"github.com/intrigues/golang-starter-boilerplate/internal/forms"
	"github.com/intrigues/golang-starter-boilerplate/internal/helpers"
	"github.com/intrigues/golang-starter-boilerplate/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func HandleGet(render core.Render) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		render.RenderTemplate(rw, r, "login.page.tmpl", &models.TemplateData{Form: forms.New(nil)})
	}
}

func HandlePost(render core.Render, logger core.Clog, cfg *config.Config, db *db.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_ = cfg.Session.RenewToken(r.Context())

		err := r.ParseForm()
		if err != nil {
			helpers.ServerError(rw, err)
			return
		}

		email := r.Form.Get("emailField")
		password := r.Form.Get("passwordField")

		form := forms.New(r.PostForm)
		form.Required("emailField", "passwordField")
		form.IsEmail("emailField")

		if !form.Valid() {
			render.RenderTemplate(rw, r, "login.page.tmpl", &models.TemplateData{
				Form: form,
			})
			return
		}

		var user models.Users
		db.Conn.First(&user, "email = ?", email)

		// authenticate method
		// TODO: isolate this method in models/controllers
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			cfg.Session.Put(r.Context(), "error", "Invalid login credentials")
			http.Redirect(rw, r, "/login", http.StatusSeeOther)
			return
		}

		cfg.Session.Put(r.Context(), "currentuser", user)
		cfg.Session.Put(r.Context(), "flash", "Logged in successfully")
		logger.Debugf("Logged in successfully. username:", user.Username)
		http.Redirect(rw, r, "/admin/dashboard", http.StatusSeeOther)
	}
}

func GetLogout(logger core.Clog, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("user logged out")
		_ = cfg.Session.Destroy(r.Context())
		_ = cfg.Session.RenewToken(r.Context())
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
