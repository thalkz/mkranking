package web

import (
	"fmt"
	"net/http"

	"github.com/thalkz/kart/internal/database"
	"github.com/thalkz/kart/internal/models"
)

type RankingPage struct {
	Season   int
	DaysLeft int
	Players  []models.Player
}

func RankingHandler(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return nil
	}

	players, err := database.GetAllPlayers()
	if err != nil {
		return fmt.Errorf("failed to get all players: %w", err)
	}

	data := RankingPage{
		Season:   1,
		DaysLeft: 60,
		Players:  players,
	}
	renderTemplate(w, "ranking.html", data)
	return nil
}
