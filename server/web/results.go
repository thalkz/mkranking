package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
)

type ResultsPage struct {
	Race         *models.Race
	ShowOkButton bool
}

func ResultsPageHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	raceIdStr := r.FormValue("race_id")
	raceId, err := strconv.Atoi(raceIdStr)
	if err != nil {
		return fmt.Errorf("failed to parse key %v: %w", raceIdStr, err)
	}

	showOkButton := r.FormValue("show_ok_button")

	race, err := database.GetRace(raceId)
	if err != nil {
		return fmt.Errorf("failed to get race details: %w", err)
	}

	data := &ResultsPage{
		Race:         race,
		ShowOkButton: showOkButton == "true",
	}
	renderTemplate(w, "results.html", data)
	return nil
}
