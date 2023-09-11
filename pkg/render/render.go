package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders templates using html
func RenderTemplate(writter http.ResponseWriter, html string) {
	parsedTemplate, err1 := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	err2 := parsedTemplate.Execute(writter, nil)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
}
