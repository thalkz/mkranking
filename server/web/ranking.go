package web

import (
	"fmt"
	"net/http"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
)

type RankingPage struct {
	Season             int
	TimeUntilSeasonEnd string
	MinRacesCount      int
	RankedPlayers      []models.Player
	UnrankedPlayers    []models.Player
}

func RankingHandler(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return nil
	}

	rankedPlayers, err := database.GetRankedPlayers(config.MinRacesCount)
	if err != nil {
		return fmt.Errorf("failed to get ranked players: %w", err)
	}

	unrankedPlayers, err := database.GetUnrankedPlayers(config.MinRacesCount)
	if err != nil {
		return fmt.Errorf("failed to get unranked players: %w", err)
	}

	timeUntil, err := parseTimeUntil("2006-01-02", config.SeasonEndDate)
	if err != nil {
		return fmt.Errorf("failed to parse timeuntil: %w", err)
	}

	data := RankingPage{
		Season:             config.Season,
		TimeUntilSeasonEnd: timeUntil,
		MinRacesCount:      config.MinRacesCount,
		RankedPlayers:      rankedPlayers,
		UnrankedPlayers:    unrankedPlayers,
	}
	renderTemplate(w, "ranking.html", data)
	return nil
}
