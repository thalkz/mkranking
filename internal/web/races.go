package web

import (
	"html/template"
	"log"
	"net/http"
)

type RacesPage struct {
	Title string
	Body  string
}

func RacesHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/races.html")
	if err != nil {
		// log.Printf("parsing template failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := RankingPage{
		Title: "Hello",
		Body:  "This is the body",
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("executing template failed: %v", err)
	}
}
