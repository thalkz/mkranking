package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type updatePlayerRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func UpdatePlayer(w http.ResponseWriter, req *http.Request) {
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
	fmt.Printf("Updating %v...\n", body.Name)
	// TODO Create player
}
