package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8082"

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "this is the about page")
	_, _ = fmt.Println(w, fmt.Sprintf("result is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
