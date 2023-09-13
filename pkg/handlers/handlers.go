package handlers

import (
	"net/http"
	"webshop/pkg/config"
	"webshop/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(writter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writter, "home.page.html")
}

// About is the about page handler
func (m *Repository) About(writter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writter, "about.page.html")
}
