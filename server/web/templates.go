package web

import (
	"html/template"
	"log"
	"net/http"
)

var t = template.Must(template.ParseGlob("../templates/*.html"))

func renderTemplate(w http.ResponseWriter, template string, data any) {
	err := t.ExecuteTemplate(w, template, data)
	if err != nil {
		log.Printf("executing template failed: %v\n", err)
	}
}
