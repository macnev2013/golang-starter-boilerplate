package handlers

import (
	"net/http"

	"github.com/intrigues/golang-starter-boilerplate/internal/config"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
