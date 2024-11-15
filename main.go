package main

import (
	"fmt"
	"net/http"
)

const portNumber = "8080"

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is the home page")
	if err != nil {
		fmt.Println("Error writing response in home:", err)
	}
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, err := fmt.Fprintf(w, "This is the about page and the sum  is %d", sum)
	if err != nil {
		fmt.Println("Error writing response in home:", err)
	}
}

func addValues(x, y int) int {
	return x + y
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
