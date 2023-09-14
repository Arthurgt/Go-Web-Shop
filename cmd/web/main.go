package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webshop/pkg/config"
	"webshop/pkg/handlers"
	"webshop/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

// main is the main application function
func main() {
	app.InProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 10 * time.Minute
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	handlersRepo := handlers.NewRepo(&app)
	handlers.NewHandlers(handlersRepo)
	render.NewTemplates(&app)

	fmt.Println("Starting application on port ", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
