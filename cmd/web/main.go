package main

import (
	"fmt"
	"github/mahdikaseatashin/wa-hello-world/pkg/config"
	"github/mahdikaseatashin/wa-hello-world/pkg/handlers"
	"github/mahdikaseatashin/wa-hello-world/pkg/render"
	"log"
	"net/http"
)

// main is the main entry point for the application
//
// The function registers the handlers for the "/" and "/about" routes
// and starts the server on port portNumber.
func main() {
	var app config.AppConfig

	app.Port = "8001"

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache: ", err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", app.Port))

	srv := &http.Server{
		Addr:    ":" + app.Port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
