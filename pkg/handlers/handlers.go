package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

	render.RenderTemplate(write, request, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m Repository) About(write http.ResponseWriter, request *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(write, request, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m Repository) Reservation(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, request, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders the room page
func (m Repository) Generals(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, request, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m Repository) Majors(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, request, "majors.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page and displays form
func (m Repository) Contact(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, request, "contact.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page and displays form
func (m Repository) Availability(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(write, request, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page and displays form
func (m Repository) PostAvailability(write http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	write.Write([]byte(fmt.Sprintf("Posted to search availability from %s to %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Mensage string `json:"mesage"`
}

// AvailabilityJSON handles request for availability and send JSON response
func (m Repository) AvailabilityJSON(write http.ResponseWriter, request *http.Request) {
	response := jsonResponse{
		Ok:      false,
		Mensage: "Not available!",
	}

	out, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Println(err)
	}

	write.Header().Set("Content-Type", "application/json")
	write.Write(out)
}
