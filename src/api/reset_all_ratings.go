package api

import (
	"fmt"
	"net/http"

	"github.com/thalkz/kart/src/database"
)

func ResetAllRatings(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Reseting all ratings...")

	players, err := database.GetAllPlayers()
	if err != nil {
		handleError(w, err)
		return
	}

	// TODO Update all ratings in the same transaction
	for i := range players {
		err := database.UpdatePlayerRating(players[i].Id, 1000.0)
		if err != nil {
			handleError(w, err)
			return
		}
	}
	fmt.Fprint(w, "ok")
}
