package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(writter http.ResponseWriter, request *http.Request) {
	renderTemplate(writter, "home.page.html")
}

// About is the about page handler
func About(writter http.ResponseWriter, request *http.Request) {
	renderTemplate(writter, "about.page.html")
}

func renderTemplate(writter http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(writter, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
