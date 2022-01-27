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

func SubmitResults(w http.ResponseWriter, req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var body submitResultsRequest
	if err := json.Unmarshal(b, &body); err != nil {
		return err
	}
	fmt.Printf("Submitting results %v\n", body.Ranking)

	if err := database.CreateRace(body.Ranking); err != nil {
		return err
	}

	players, err := database.GetPlayers(body.Ranking)
	if err != nil {
		return err
	}

	ratings := make([]float64, len(players))
	for i := range players {
		ratings[i] = players[i].Rating
	}

	newRatings := elo.ComputeRatings(ratings)

	// TODO Update all ratings in the same transaction
	for i := range body.Ranking {
		if err := database.UpdatePlayerRating(body.Ranking[i], newRatings[i]); err != nil {
			return err
		}
	}
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
