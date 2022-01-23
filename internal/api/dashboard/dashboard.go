package dashboard

import (
	"net/http"

	"github.com/intrigues/golang-starter-boilerplate/core"
	"github.com/intrigues/golang-starter-boilerplate/internal/forms"
	"github.com/intrigues/golang-starter-boilerplate/internal/models"
)

func HandleGet(render core.Render) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		render.RenderTemplate(rw, r, "dashboard.page.tmpl", &models.TemplateData{Form: forms.New(nil)})
	}
}
