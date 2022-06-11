package web

import (
	"net/http"
)

func PlayerHandler(w http.ResponseWriter, r *http.Request) error {
	renderTemplate(w, "player.html", nil)
	return nil
}
