package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
	"github.com/thalkz/kart/utils"
)

func HistoryHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	season := utils.ParseSeason(cfg, r)

	players, err := database.GetRankedPlayers(season, cfg.MinRacesCount)
	if err != nil {
		return fmt.Errorf("failed to get all players: %w", err)
	}

	events, err := database.GetRankedPlayersEvents(cfg, season)
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
