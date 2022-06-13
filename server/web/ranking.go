package web

import (
	"fmt"
	"net/http"
	"strconv"

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

func parseSeason(r *http.Request) int {
	seasonStr := r.FormValue("season")
	season, err := strconv.Atoi(seasonStr)
	if seasonStr == "" || err != nil {
		return config.Season
	}
	return season
}

func RankingHandler(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return nil
	}

	season := parseSeason(r)

	rankedPlayers, err := database.GetRankedPlayers(season, config.MinRacesCount)
	if err != nil {
		return fmt.Errorf("failed to get ranked players: %w", err)
	}

	unrankedPlayers, err := database.GetUnrankedPlayers(season, config.MinRacesCount)
	if err != nil {
		return fmt.Errorf("failed to get unranked players: %w", err)
	}

	timeUntil, err := parseTimeUntil("2006-01-02", config.SeasonEndDate)
	if err != nil {
		return fmt.Errorf("failed to parse timeuntil: %w", err)
	}

	data := RankingPage{
		Season:             season,
		TimeUntilSeasonEnd: timeUntil,
		MinRacesCount:      config.MinRacesCount,
		RankedPlayers:      rankedPlayers,
		UnrankedPlayers:    unrankedPlayers,
	}
	renderTemplate(w, "ranking.html", data)
	return nil
}
