package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templateCache = make(map[string]*template.Template)

// RenderTemplate renders templates using html
func RenderTemplate(writter http.ResponseWriter, html string) {
	var tmpl *template.Template
	var err error

	_, inMap := templateCache[html]
	if !inMap {
		log.Println("Creating template and adding to cache ", html)
		err = createTemplateCahce(html)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	log.Println("Using cached template ", html)
	tmpl = templateCache[html]
	err = tmpl.Execute(writter, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// createTemplateCahce creates cache to store already readed files
func createTemplateCahce(html string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", html),
		"./templates/base.layout.html",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	templateCache[html] = tmpl
	return nil
}
