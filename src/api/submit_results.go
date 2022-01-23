package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type submitResultsRequest struct {
	Results []string `json:"results"`
}

type submitResultsResponse struct {
	Players []player `json:"players"`
}

func SubmitResults(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	var body submitResultsRequest
	jsonErr := json.Unmarshal(b, &body)
	if jsonErr != nil {
		handleError(w, jsonErr)
		return
	}

	fmt.Printf("Submitting results %v...\n", body.Results)
	submitResultsResponse := submitResultsResponse{
		Players: []player{
			{
				Id:     "one",
				Name:   "One",
				Rating: 1000.0,
			},
			{
				Id:     "two",
				Name:   "Two",
				Rating: 1002.0,
			},
		},
	}

	responseBytes, resErr := json.Marshal(&submitResultsResponse)
	if resErr != nil {
		handleError(w, resErr)
		return
	}
	fmt.Fprint(w, string(responseBytes))
}
