package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github/mahdikaseatashin/wa-hello-world/pkg/config"
	"github/mahdikaseatashin/wa-hello-world/pkg/handlers"
	"github/mahdikaseatashin/wa-hello-world/pkg/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig

const portNumber = "8001"

var session *scs.SessionManager

// main is the main entry point for the application
//
// The function registers the handlers for the "/" and "/about" routes
// and starts the server on port portNumber.
func main() {
	app.InProduction = false
	app.Port = portNumber

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
