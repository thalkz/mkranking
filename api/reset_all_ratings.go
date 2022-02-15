package api

import (
	"encoding/json"
	"net/http"

	"github.com/thalkz/kart/database"
)

func ResetAllRatings(w http.ResponseWriter, req *http.Request) error {
	players, err := database.GetAllPlayers()
	if err != nil {
		return err
	}

	ids := make([]int, len(players))
	ratings := make([]float64, len(players))
	for i := range players {
		ids[i] = players[i].Id
		ratings[i] = 1000.0
	}

	if err = database.UpdatePlayerRatings(ids, ratings); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
