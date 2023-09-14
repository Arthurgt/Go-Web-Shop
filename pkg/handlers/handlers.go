package handlers

import (
	"net/http"
	"webshop/pkg/config"
	"webshop/pkg/models"
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
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(writter, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(writter http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."
	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(writter, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
