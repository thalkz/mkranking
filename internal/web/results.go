package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/internal/database"
)

func ResultsPageHandler(w http.ResponseWriter, r *http.Request) error {
	raceIdStr := r.FormValue("race_id")
	raceId, err := strconv.Atoi(raceIdStr)
	if err != nil {
		return fmt.Errorf("failed to parse key %v: %w", raceIdStr, err)
	}

	raceDetails, err := database.GetRaceDetails(raceId)
	if err != nil {
		return fmt.Errorf("failed to get race details: %w", err)
	}

	renderTemplate(w, "results.html", raceDetails)
	return nil
}
