package core

import (
	"net/http"

	"github.com/intrigues/golang-starter-boilerplate/internal/models"
)

type Render interface {
	RenderTemplate(http.ResponseWriter, *http.Request, string, *models.TemplateData) error
}
