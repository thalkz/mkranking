package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type deletePlayerRequest struct {
	Id string `json:"id"`
}

func DeletePlayer(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	var body deletePlayerRequest
	jsonErr := json.Unmarshal(b, &body)
	if jsonErr != nil {
		handleError(w, jsonErr)
		return
	}
	fmt.Fprintf(w, "Deleting %v...\n", body.Id)
	// TODO Delete player
}
