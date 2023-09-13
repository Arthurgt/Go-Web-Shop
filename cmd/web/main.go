package main

import (
	"fmt"
	"log"
	"net/http"
	"webshop/pkg/config"
	"webshop/pkg/handlers"
	"webshop/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	handlersRepo := handlers.NewRepo(&app)
	handlers.NewHandlers(handlersRepo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
