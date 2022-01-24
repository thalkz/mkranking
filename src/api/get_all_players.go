package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thalkz/kart/src/database"
)

func GetAllPlayers(w http.ResponseWriter, req *http.Request) {
	players, err := database.GetAllPlayers()
	if err != nil {
		handleError(w, err)
		return
	}
	bytes, err := json.Marshal(players)
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Fprintf(w, "%v", string(bytes))
}