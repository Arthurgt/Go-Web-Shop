package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(writter http.ResponseWriter, html string) {
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	template, ok := templateCache[html]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = template.Execute(buf, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = buf.WriteTo(writter)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}
	}
	return myCache, nil
}
