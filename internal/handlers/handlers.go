package handlers

import (
	"github.com/Animenkor/AthenaWebApp/internal/config"
	"github.com/Animenkor/AthenaWebApp/internal/models"
	"github.com/Animenkor/AthenaWebApp/internal/render"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
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

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Tasks(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "tasks.page.tmpl", &models.TemplateData{})
}
