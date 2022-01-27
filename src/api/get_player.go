package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/thalkz/kart/src/database"
	"github.com/thalkz/kart/src/models"
)

type getPlayerRequest struct {
	Id int `json:"id"`
}

func GetPlayer(w http.ResponseWriter, req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var body getPlayerRequest
	if err = json.Unmarshal(b, &body); err != nil {
		return err
	}

	var player models.Player
	player, err = database.GetPlayer(body.Id)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
		Data:   player,
	})
}
