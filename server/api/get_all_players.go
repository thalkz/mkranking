package api

import (
	"encoding/json"
	"net/http"

	"github.com/thalkz/kart/database"
)

func GetAllPlayers(w http.ResponseWriter, req *http.Request) error {
	players, err := database.GetAllPlayers()
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
		Data:   players,
	})
}
