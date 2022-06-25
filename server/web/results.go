package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/database"
)

func ResultsPageHandler(w http.ResponseWriter, r *http.Request) error {
	raceIdStr := r.FormValue("race_id")
	raceId, err := strconv.Atoi(raceIdStr)
	if err != nil {
		return fmt.Errorf("failed to parse key %v: %w", raceIdStr, err)
	}

	race, err := database.GetRace(raceId)
	if err != nil {
		return fmt.Errorf("failed to get race details: %w", err)
	}

	renderTemplate(w, "results.html", race)
	return nil
}