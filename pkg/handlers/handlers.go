package handlers

import (
	"net/http"

	"github.com/rudsonalves/webapp/pkg/config"
	"github.com/rudsonalves/webapp/pkg/models"
	"github.com/rudsonalves/webapp/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Repo the repository used by the handlers
var Repo *Repository

// Home is the home page handler
func (m Repository) Home(write http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(write, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m Repository) About(write http.ResponseWriter, req *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(write, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
