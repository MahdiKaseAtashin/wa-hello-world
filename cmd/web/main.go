package main

import (
	"fmt"
	"github/mahdikaseatashin/wa-hello-world/pkg/handlers"
	"net/http"
)

const portNumber = "8080"

// main is the main entry point for the application
//
// The function registers the handlers for the "/" and "/about" routes
// and starts the server on port portNumber.
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err := http.ListenAndServe(":"+portNumber, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
