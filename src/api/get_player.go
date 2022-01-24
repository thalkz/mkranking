package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thalkz/kart/src/database"
	"github.com/thalkz/kart/src/models"
)

type getPlayerRequest struct {
	Id int `json:"id"`
}

func GetPlayer(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(w, err)
		return
	}

	var body getPlayerRequest
	err = json.Unmarshal(b, &body)
	if err != nil {
		handleError(w, err)
		return
	}

	var player models.Player
	player, err = database.GetPlayer(body.Id)
	if err != nil {
		handleError(w, err)
		return
	}

	var bytes []byte
	bytes, err = json.Marshal(player)
	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Fprintf(w, string(bytes))
}
