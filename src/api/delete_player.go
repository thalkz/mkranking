package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/thalkz/kart/src/database"
)

type deletePlayerRequest struct {
	Id int `json:"id"`
}

func DeletePlayer(w http.ResponseWriter, req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var body deletePlayerRequest
	if err := json.Unmarshal(b, &body); err != nil {
		return err
	}

	if err := database.DeletePlayer(body.Id); err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
	})
}
