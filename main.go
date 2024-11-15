package main

import (
	"fmt"
	"net/http"
)

const portNumber = "8080"

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err := http.ListenAndServe(":"+portNumber, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
