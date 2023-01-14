package main

import (
	"fmt"
	"github.com/webhire/gowebapp/pkg/config"
	"github.com/webhire/gowebapp/pkg/handlers"
	"github.com/webhire/gowebapp/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8081"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cant create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting server on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
