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
	Status string `json:"status"`
	Id     int    `json:"id"`
}

const (
	initialRating = 1000.0
)

func CreatePlayer(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	var body createPlayerRequest
	jsonErr := json.Unmarshal(b, &body)
	if jsonErr != nil {
		handleError(w, jsonErr)
		return
	}
	fmt.Printf("Creating %v...\n", body.Name)

	id, dbErr := database.CreatePlayer(body.Name, initialRating)
	if dbErr != nil {
		handleError(w, dbErr)
		return
	}
	responseBody := createPlayerResponse{
		Status: "ok",
		Id:     id,
	}
	responseJson, _ := json.Marshal(responseBody)
	fmt.Fprintf(w, "%v", string(responseJson))
}
