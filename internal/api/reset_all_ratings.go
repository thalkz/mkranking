package api

import (
	"encoding/json"
	"net/http"

	"github.com/thalkz/kart/internal/config"
	"github.com/thalkz/kart/internal/database"
)

func ResetAllRatings(w http.ResponseWriter, req *http.Request) error {
	if err := database.ResetAllRatings(config.InitialRating); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
