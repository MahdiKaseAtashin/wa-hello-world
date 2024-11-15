package render

import (
	"html/template"
	"log"
	"net/http"
)

var tc = make(map[string]*template.Template)

func FromTemplatePath(w http.ResponseWriter, path string) error {
	var parsedTemplate *template.Template
	var err error

	// check if we already have the template in our cache
	if _, inMap := tc[path]; inMap {
		log.Println("using cached template...")
	} else {
		log.Println("creating template and adding cache...")
		err = createTemplateCache(path)
		if err != nil {
			return err
		}
	}

	parsedTemplate = tc[path]

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		return err
	}
	return nil
}

func createTemplateCache(path string) error {
	templates := []string{
		"./templates/" + path,
		"./templates/base.layout.html",
	}
	parsedTemplate, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[path] = parsedTemplate
	return nil
}
