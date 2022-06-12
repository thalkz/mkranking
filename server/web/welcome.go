package web

import (
	"net/http"
)

type welcomePlayerPage struct {
	Name string
	Icon string
}

func WelcomePlayerPage(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	icon := r.FormValue("icon")

	data := welcomePlayerPage{
		Name: name,
		Icon: icon,
	}
	renderTemplate(w, "welcome.html", data)
	return nil
}
