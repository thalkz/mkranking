package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/thalkz/kart/database"
)

type updatePlayerRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func UpdatePlayer(w http.ResponseWriter, req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var body updatePlayerRequest
	if err := json.Unmarshal(b, &body); err != nil {
		return err
	}
	log.Printf("Updating %v\n", body.Name)

	err = database.UpdatePlayerName(body.Id, body.Name)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
