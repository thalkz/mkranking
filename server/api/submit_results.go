package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	players, err := database.GetPlayers(body.Ranking)
	log.Printf("Getting players: %v\n", players)
	if err != nil {
		return err
	}

	oldRatings := make([]float64, len(players))
	for i := range players {
		oldRatings[i] = players[i].Rating
	}

	log.Printf("Computing elo with ratings: %v\n", oldRatings)
	newRatings := elo.ComputeRatings(oldRatings)

	log.Printf("New ratings are %v\n", newRatings)

	log.Printf("Creating race with ranking: %v\n", body.Ranking)
	if err := database.CreateRace(body.Ranking, oldRatings, newRatings); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
