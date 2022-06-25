package web

import (
	"fmt"
	"net/http"

	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
)

type racesPage struct {
	Races []models.Race
}

func RacesHandler(w http.ResponseWriter, r *http.Request) error {
	season := parseSeason(r)

	races, err := database.GetAllRaces(season)
	if err != nil {
		return fmt.Errorf("failed to get all races: %w", err)
	}

	// Convert to timeago
	for i := range races {
		timeago, err := parseTimeAgo("2006-01-02T15:04:05Z", races[i].Date)
		if err != nil {
			return fmt.Errorf("failed to parse timeago: %w", err)
		}
		races[i].Date = timeago
	}

	data := racesPage{
		Races: races,
	}
	renderTemplate(w, "races.html", data)
	return nil
}
