package api

import (
	"encoding/json"
	"net/http"
)

type RatingsHistory struct {
	PlayerNames []string    `json:"player_names"`
	Dates       []string    `json:"dates"`
	DataRows    [][]float64 `json:"data_rows"`
}

func GetRatingsHistory(w http.ResponseWriter, req *http.Request) error {
	// TODO Return actual data
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
		Data: &RatingsHistory{
			PlayerNames: []string{"PLayerA", "PlayerB"},
			Dates:       []string{"2000-01-01", "2000-01-02", "2000-01-03", "2000-01-04"},
			DataRows: [][]float64{
				{1000.0, 1101.0, 1001.0, 1102.0},
				{1002.0, 1103.0, 1003.0, 1104.0},
			},
		},
	})
}
