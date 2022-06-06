package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thalkz/kart/database"
)

type getPlayerRacesRequest struct {
	Id int `json:"id"`
}

func GetPlayerRaces(w http.ResponseWriter, req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	var body getPlayerRacesRequest
	if err = json.Unmarshal(b, &body); err != nil {
		return fmt.Errorf("failed to unmarshal request body: %w", err)
	}

	races, err := database.GetPlayerRaces(body.Id)
	if err != nil {
		return fmt.Errorf("failed to get player races: %w", err)
	}

	err = json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
		Data:   races,
	})
	if err != nil {
		return fmt.Errorf("failed to encore response: %w", err)
	}
	return nil
}
