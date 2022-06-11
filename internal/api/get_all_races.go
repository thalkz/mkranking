package api

import (
	"encoding/json"
	"net/http"

	"github.com/thalkz/kart/internal/database"
)

func GetAllRaces(w http.ResponseWriter, req *http.Request) error {
	races, err := database.GetAllRaces()
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
		Data:   races,
	})
}
