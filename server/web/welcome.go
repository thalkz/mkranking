package web

import (
	"net/http"

	"github.com/thalkz/kart/config"
)

type welcomePlayerPage struct {
	Name string
	Icon string
}

func WelcomePlayerPage(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	icon := r.FormValue("icon")

	data := welcomePlayerPage{
		Name: name,
		Icon: icon,
	}
	renderTemplate(w, "welcome.html", data)
	return nil
}
