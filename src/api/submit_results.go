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

	fmt.Printf("Creating race with ranking %v...\n", body.Ranking)
	if err := database.CreateRace(body.Ranking); err != nil {
		return err
	}

	fmt.Println("Getting players...")
	players, err := database.GetPlayers(body.Ranking)
	if err != nil {
		return err
	}

	ratings := make([]float64, len(players))
	for i := range players {
		ratings[i] = players[i].Rating
	}

	fmt.Println("Computing elo...")
	newRatings := elo.ComputeRatings(ratings)

	fmt.Printf("Updating ratings %v...\n", newRatings)
	if err := database.UpdatePlayerRatings(body.Ranking, newRatings); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
