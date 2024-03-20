package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(rw http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
