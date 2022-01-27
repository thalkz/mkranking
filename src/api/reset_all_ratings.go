package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thalkz/kart/src/database"
)

func ResetAllRatings(w http.ResponseWriter, req *http.Request) error {
	fmt.Printf("Reseting all ratings...")

	players, err := database.GetAllPlayers()
	if err != nil {
		return err
	}

	// TODO Update all ratings in the same transaction
	for i := range players {
		err := database.UpdatePlayerRating(players[i].Id, 1000.0)
		if err != nil {
			return err
		}
	}
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
