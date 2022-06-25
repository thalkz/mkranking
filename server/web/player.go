package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
)

type PlayerPage struct {
	Player models.Player
	Races  []models.Race
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) error {
	playerIdStr := r.FormValue("id")
	playerId, err := strconv.Atoi(playerIdStr)
	if err != nil {
		return fmt.Errorf("failed to parse key %v: %w", playerIdStr, err)
	}

	player, err := database.GetPlayer(playerId)
	if err != nil {
		return fmt.Errorf("failed to get player: %w", err)
	}

	playerRaces, err := database.GetPlayerRaces(playerId)
	if err != nil {
		return fmt.Errorf("failed to get player races: %w", err)
	}

	// Convert to timeago
	for i := range playerRaces {
		timeago, err := parseTimeAgo("2006-01-02T15:04:05Z", playerRaces[i].Date)
		if err != nil {
			return fmt.Errorf("failed to parse timeago: %w", err)
		}
		playerRaces[i].Date = timeago
	}

	data := &PlayerPage{
		Player: player,
		Races:  playerRaces,
	}

	renderTemplate(w, "player.html", data)
	return nil
}
