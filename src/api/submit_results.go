package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thalkz/kart/src/database"
	"github.com/thalkz/kart/src/elo"
)

type submitResultsRequest struct {
	Ranking []int `json:"ranking"`
}

func SubmitResults(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	var body submitResultsRequest
	jsonErr := json.Unmarshal(b, &body)
	if jsonErr != nil {
		handleError(w, jsonErr)
		return
	}

	fmt.Printf("Submitting results %v\n", body.Ranking)

	players, err := database.GetPlayers(body.Ranking)
	if err != nil {
		handleError(w, err)
		return
	}

	ratings := make([]float64, len(players))
	for i := range players {
		ratings[i] = players[i].Rating
	}

	newRatings := elo.ComputeRatings(ratings)

	// TODO Update all ratings in the same transaction
	for i := range body.Ranking {
		err := database.UpdatePlayerRating(body.Ranking[i], newRatings[i])
		if err != nil {
			handleError(w, err)
			return
		}
	}
	fmt.Fprint(w, "ok")
}
