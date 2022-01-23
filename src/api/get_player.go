package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type getPlayerRequest struct {
	Id string `json:"id"`
}

func GetPlayer(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	var body getPlayerRequest
	jsonErr := json.Unmarshal(b, &body)
	if jsonErr != nil {
		handleError(w, jsonErr)
		return
	}
	fmt.Printf("Getting \"%v\"...\n", body.Id)
	// TODO Create player
}
