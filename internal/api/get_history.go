package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thalkz/kart/internal/database"
)

type getHistoryRequest struct {
	PlayerIds []int `json:"ids"`
}

type getHistoryResponse struct {
	PlayerIds []int       `json:"player_ids"`
	Dates     []string    `json:"dates"`
	DataRows  [][]float64 `json:"data_rows"`
}

func GetHistory(w http.ResponseWriter, req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	var body getHistoryRequest
	if err := json.Unmarshal(b, &body); err != nil {
		return fmt.Errorf("failed to unmarshal body: %w", err)
	}

	if len(body.PlayerIds) == 0 {
		return fmt.Errorf("should specify at least 1 player id, got 0")
	}

	dataRows := make([][]float64, 0)
	var dates []string
	for _, userId := range body.PlayerIds {
		userHistory, err := database.GetPlayerHistory(userId)
		if err != nil {
			return fmt.Errorf("failed to get history for %v: %w", userId, err)
		}

		// Add dates if missing
		if len(dates) == 0 {
			dates = make([]string, len(userHistory))
			for i, event := range userHistory {
				dates[i] = event.Date
			}
		}

		// Add row
		var row = make([]float64, len(userHistory))
		for i, event := range userHistory {
			row[i] = event.NewRating
		}
		dataRows = append(dataRows, row)
	}

	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
		Data: &getHistoryResponse{
			PlayerIds: body.PlayerIds,
			Dates:     dates,
			DataRows:  dataRows,
		},
	})
}
