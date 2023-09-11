package handlers

import (
	"net/http"
	"webshop/pkg/render"
)

// Home is the home page handler
func Home(writter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writter, "home.page.html")
}

// About is the about page handler
func About(writter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writter, "about.page.html")
}
