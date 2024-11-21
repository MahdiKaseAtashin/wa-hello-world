package main

import (
	"github.com/gorilla/mux"
	"github/mahdikaseatashin/wa-hello-world/pkg/config"
	"github/mahdikaseatashin/wa-hello-world/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Repo.Home)
	r.HandleFunc("/about", handlers.Repo.About)

	return r
}
