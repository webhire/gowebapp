package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8084"

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

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100, 10)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100, 10, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
