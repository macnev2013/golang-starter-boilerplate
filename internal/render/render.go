package render

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"

	"github.com/intrigues/golang-starter-boilerplate/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

func (render *render) addDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = render.cfg.Session.PopString(r.Context(), "flash")
	td.Error = render.cfg.Session.PopString(r.Context(), "error")
	td.Warning = render.cfg.Session.PopString(r.Context(), "warning")
	td.CurrentUser = render.cfg.Session.Get(r.Context(), "currentuser")
	td.CurrentPage = r.RequestURI
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders a template
func (render *render) RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc TemplateCacheType

	if render.cfg.Env == "prod" {
		tc = render.templateCache
	} else {
		tc, _ = GenearteTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		render.logger.Errorf("error retriving template cache")
		return errors.New("can't get template from cache")
	}
	buf := new(bytes.Buffer)
	td = render.addDefaultData(td, r)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		render.logger.Errorf("error writing template to browser", err)
		return err
	}
	return nil
}
