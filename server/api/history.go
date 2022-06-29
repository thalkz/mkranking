package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request) error {
	players, err := database.GetAllPlayers(2)
	if err != nil {
		return fmt.Errorf("failed to get all players: %w", err)
	}

	events, err := database.GetAllPlayersEvents(2) // TODO Replace with season param
	if err != nil {
		return fmt.Errorf("failed to get all history: %w", err)
	}

	if len(players) != len(events) {
		return fmt.Errorf("players and events should have the same length: found %v and %v", len(players), len(events))
	}

	response := &models.History{
		Players: players,
		Events:  events,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
}
