package web

import (
	"net/http"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) error {
	renderTemplate(w, "stats.html", nil)
	return nil
}
