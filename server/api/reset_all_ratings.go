package api

import (
	"encoding/json"
	"net/http"

	"github.com/thalkz/kart/database"
)

func ResetAllRatings(w http.ResponseWriter, req *http.Request) error {
	if err := database.ResetAllRatings(initialRating); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
