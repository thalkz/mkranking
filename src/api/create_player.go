package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thalkz/kart/src/database"
)

type createPlayerRequest struct {
	Name string `json:"name"`
}

type createPlayerResponse struct {
	Id int `json:"id"`
}

const (
	initialRating = 1000.0
)

func CreatePlayer(w http.ResponseWriter, req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	var body createPlayerRequest
	err = json.Unmarshal(b, &body)
	if err != nil {
		return err
	}
	fmt.Printf("Creating %v\n", body.Name)

	id, err := database.CreatePlayer(body.Name, initialRating)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(
		&JsonResponse{
			Status: "ok",
			Data: &createPlayerResponse{
				Id: id,
			},
		},
	)
}
