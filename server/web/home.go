package web

import (
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Title string
	Body  string
}

func Home(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/edit/"):]
	// log.Println(title)

	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Printf("parsing template failed: %v", err)
	}

	data := HomePage{
		Title: "Hello",
		Body:  "This is the body",
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("executing template failed: %v", err)
	}
}
