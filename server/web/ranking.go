package web

import (
	"fmt"
	"net/http"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
	"github.com/thalkz/kart/utils"
)

type RankingPage struct {
	Season              int
	TimeUntilEnd        string
	IsCompetitionActive bool
	MinRacesCount       int
	RankedPlayers       []models.Player
	UnrankedPlayers     []models.Player
}

func RankingHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return nil
	}

	season := utils.ParseSeason(cfg, r)

	rankedPlayers, err := database.GetRankedPlayers(season, cfg.MinRacesCount)
	if err != nil {
		return fmt.Errorf("failed to get ranked players: %w", err)
	}

	unrankedPlayers, err := database.GetUnrankedPlayers(season, cfg.MinRacesCount)
	if err != nil {
		return fmt.Errorf("failed to get unranked players: %w", err)
	}

	isCompetitionActive := cfg.IsCompetitionActive()
	var timeUntilEnd string

	if isCompetitionActive {
		timeUntilEnd, err = utils.ParseTimeUntil(cfg.GetCompetitionEndDate())
		if err != nil {
			return fmt.Errorf("failed to parse timeuntil: %w", err)
		}
	} else {
		timeUntilEnd, err = utils.ParseTimeUntil(cfg.GetNextSeasonStartDate())
		if err != nil {
			return fmt.Errorf("failed to parse timeuntil: %w", err)
		}
	}

	data := RankingPage{
		Season:              season,
		TimeUntilEnd:        timeUntilEnd,
		IsCompetitionActive: isCompetitionActive,
		MinRacesCount:       cfg.MinRacesCount,
		RankedPlayers:       rankedPlayers,
		UnrankedPlayers:     unrankedPlayers,
	}
	renderTemplate(w, "ranking.html", data)
	return nil
}
