package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, world")
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	})

	err := http.ListenAndServe(":7551", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
