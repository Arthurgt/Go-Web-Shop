package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(writter http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(writter, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
