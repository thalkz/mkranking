package web

import (
	"net/http"

	"github.com/thalkz/kart/config"
)

func StatsHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	renderTemplate(w, "stats.html", nil)
	return nil
}
