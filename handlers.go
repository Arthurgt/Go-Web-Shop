package main

import (
	"net/http"
)

// Home is the home page handler
func Home(writter http.ResponseWriter, request *http.Request) {
	renderTemplate(writter, "home.page.html")
}

// About is the about page handler
func About(writter http.ResponseWriter, request *http.Request) {
	renderTemplate(writter, "about.page.html")
}
