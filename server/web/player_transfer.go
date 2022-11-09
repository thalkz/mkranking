package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
)

type chooseSyncPlayerPage struct {
	LastPlayers []models.Player
}

func TransferPlayerHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	playerIdStr := r.FormValue("id")
	if playerIdStr != "" {
		return syncPlayerHandler(cfg, w, r, playerIdStr)
	} else {
		return choosePlayerSyncHandler(cfg, w, r)
	}
}

func syncPlayerHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request, playerIdStr string) error {
	oldPlayerId, err := strconv.Atoi(playerIdStr)
	if err != nil {
		return fmt.Errorf("failed to parse playerId: %w", err)
	}

	err = database.TransferPlayer(oldPlayerId, cfg.GetSeason())
	if err != nil {
		return fmt.Errorf("failed to transfer player: %w", err)
	}
	log.Printf("Transfered player with id=%v\n", oldPlayerId)

	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}

func choosePlayerSyncHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	lastSeason := cfg.GetSeason() - 1
	players, err := database.GetAllPlayers(lastSeason)
	if err != nil {
		return fmt.Errorf("failed to get all players: %w", err)
	}

	if len(players) == 0 {
		http.Redirect(w, r, "/new", http.StatusFound)
		return nil
	}

	data := chooseSyncPlayerPage{
		LastPlayers: players,
	}
	renderTemplate(w, "player_transfer.html", data)
	return nil
}
