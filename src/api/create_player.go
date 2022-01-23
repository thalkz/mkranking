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
	Id string `json:"id"`
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
	fmt.Printf("Creating %v...\n", body.Name)

	insertStmt := `insert into "Players"("name", "rating") values('John', 1000.0)`
	result, dbErr := database.Exec(insertStmt)
	if dbErr != nil {
		handleError(w, dbErr)
	}

	fmt.Printf("Database response: %v", result)
}
