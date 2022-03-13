package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/elo"
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

	if len(body.Ranking) < 2 {
		return fmt.Errorf("cannot submit results with less than two players. recieved %v", body.Ranking)
	}

	fmt.Printf("Creating race with ranking: %v\n", body.Ranking)
	if err := database.CreateRace(body.Ranking); err != nil {
		return err
	}

	players, err := database.GetPlayers(body.Ranking)
	fmt.Printf("Getting players: %v\n", players)
	if err != nil {
		return err
	}

	ratings := make([]float64, len(players))
	for i := range players {
		ratings[i] = players[i].Rating
	}

	fmt.Printf("Computing elo with ratings: %v\n", ratings)
	newRatings := elo.ComputeRatings(ratings)

	fmt.Printf("New ratings are %v\n", newRatings)
	if err := database.UpdatePlayerRatings(body.Ranking, newRatings); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
