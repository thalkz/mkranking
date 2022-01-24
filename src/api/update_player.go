package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thalkz/kart/src/database"
)

type updatePlayerRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func UpdatePlayer(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	var body updatePlayerRequest
	jsonErr := json.Unmarshal(b, &body)
	if jsonErr != nil {
		handleError(w, jsonErr)
		return
	}
	fmt.Printf("Updating %v...\n", body.Name)

	err = database.UpdatePlayerName(body.Id, body.Name)
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Fprintf(w, "ok")
}
