package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders templates using html
func RenderTemplate(writter http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(writter, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
