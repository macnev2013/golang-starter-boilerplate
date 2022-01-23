package signup

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
		render.RenderTemplate(rw, r, "signup.page.tmpl", &models.TemplateData{
			Form: forms.New(nil),
		})
	}
}

func HandlePost(render core.Render, logger core.Clog, cfg *config.Config, db *db.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			helpers.ServerError(rw, err)
			return
		}
		logger.Debugf("form content: %+v", r.Form)

		form := forms.New(r.PostForm)
		form.Required("usernameField", "emailField", "passwordField", "roleSelector")
		form.IsEmail("emailField")
		if !form.Valid() {
			cfg.Session.Put(r.Context(), "error", "Please enter valid details")
			logger.Errorf("error validating form %v", err)
			render.RenderTemplate(rw, r, "signup.page.tmpl", &models.TemplateData{
				Form: form,
			})
			return
		}

		password_hash, err := bcrypt.GenerateFromPassword([]byte(r.Form.Get("passwordField")), 0)
		if err != nil {
			logger.Errorf("error generating hash from given password")
		}
		db.Conn.Create(&models.Users{
			Username:          r.Form.Get("usernameField"),
			Email:             r.Form.Get("emailField"),
			Password:          string(password_hash),
			IncorrectPassword: 0,
			Status:            0,
			Role:              r.Form.Get("roleSelector"),
		})

		http.Redirect(rw, r, "/login", http.StatusSeeOther)
	}
}
