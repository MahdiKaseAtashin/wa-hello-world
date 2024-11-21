package handlers

import (
	"github/mahdikaseatashin/wa-hello-world/pkg/config"
	"github/mahdikaseatashin/wa-hello-world/pkg/models"
	"github/mahdikaseatashin/wa-hello-world/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository instance with the provided AppConfig.
// It returns a pointer to the newly created Repository.
// The AppConfig is used to initialize the App field in the Repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers creates a new handlers and sets the repository used by the
// handlers. This is the entry point for the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (r *Repository) Home(w http.ResponseWriter, hr *http.Request) {
	remoteIP := hr.RemoteAddr
	r.App.Session.Put(hr.Context(), "remote_ip", remoteIP)
	err := render.FromTemplatePath(w, "home.page.html", &models.TemplateData{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// About is the about page handler
func (r *Repository) About(w http.ResponseWriter, hr *http.Request) {
	remoteIP := r.App.Session.GetString(hr.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	stringMap["remote_ip"] = remoteIP

	err := render.FromTemplatePath(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
