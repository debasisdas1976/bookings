package handlers

import (
	"fmt"
	"net/http"

	"github.com/debasisdas1976/bookings/pkg/config"
	"github.com/debasisdas1976/bookings/pkg/models"
	"github.com/debasisdas1976/bookings/pkg/render"
)

// TemplateData holds data sent from handlers to templates

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo create a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler is getting invoked")
	remoteIP := r.RemoteAddr
	fmt.Println(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello Again"

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) ContactUs(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contactus.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Orphan(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "orphan.page.tmpl", &models.TemplateData{})
}
