package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type createPlayerRequest struct {
	Name string `json:"name"`
}

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
	fmt.Fprintf(w, "Creating \"%v\"...\n", body.Name)
	// TODO Create player
}
