package handlers

import (
	"github/mahdikaseatashin/wa-hello-world/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	err := render.FromTemplatePath(w, "home.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	err := render.FromTemplatePath(w, "about.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
