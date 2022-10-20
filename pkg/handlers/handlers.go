package handlers

import (
	"net/http"

	"github.com/rudsonalves/bookings/pkg/config"
	"github.com/rudsonalves/bookings/pkg/models"
	"github.com/rudsonalves/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

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

// Home is the home page handler
func (m Repository) Home(write http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(write, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m Repository) About(write http.ResponseWriter, request *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(write, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m Repository) Reservation(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders the room page
func (m Repository) Generals(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m Repository) Majors(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, "majors.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page and displays form
func (m Repository) Contact(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, "contact.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page and displays form
func (m Repository) Availability(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, "search-availability.page.tmpl", &models.TemplateData{})
}
