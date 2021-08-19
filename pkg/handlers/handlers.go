package handlers

import (
	"BookingsApp/pkg/config"
	"BookingsApp/pkg/models"
	"BookingsApp/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}

}

//NewHandlers sets the Repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_IP", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform logic operation
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_IP")
	stringMap["remote_ip"] = remoteIP
	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
