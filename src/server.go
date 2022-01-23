package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

type createPlayerRequest struct {
	Name string `json:"name"`
}

func createPlayer(w http.ResponseWriter, req *http.Request) {
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

type submitResultsRequest struct {
	Results []string `json:"results"`
}

type submitResultsResponse struct {
	Players []player `json:"players"`
}

type player struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

func handleError(w http.ResponseWriter, err error) {
	fmt.Printf("Error: %v\n", err)
	fmt.Fprintf(w, "Error: %v", err)
}

func submitResults(w http.ResponseWriter, req *http.Request) {
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

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/createPlayer", createPlayer)
	http.HandleFunc("/submitResults", submitResults)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	fmt.Println("Listening on port", httpPort)
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		fmt.Println("err")
	}
}
