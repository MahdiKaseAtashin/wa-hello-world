package main

import (
	"errors"
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
		fmt.Println("Error writing response in about:", err)
	}
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x, y := 10, 2
	f, err := divideValues(float32(x), float32(y))
	if err != nil {
		fmt.Println("Error writing response in divide:", err)
		return
	}
	_, err = fmt.Fprintf(w, "This is the divide page and the %d divided by %d is %f", x, y, f)
	if err != nil {
		fmt.Println("Error writing response in divide:", err)
	}
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	return x / y, nil
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err := http.ListenAndServe(":"+portNumber, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
