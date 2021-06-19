//Handler package functions
package handlers

import (
	"net/http"

	"github.com/t-Ikonen/hellowebapp/pkg/config"
	"github.com/t-Ikonen/hellowebapp/pkg/models"
	"github.com/t-Ikonen/hellowebapp/pkg/render"
)

// Repo used by handlers
var Repo *Repository

//Repository is repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the Handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home fnction hadles Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to home pagae")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	//send data to template
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About func handles About page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
