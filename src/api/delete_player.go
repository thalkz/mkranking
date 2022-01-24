package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thalkz/kart/src/database"
)

type deletePlayerRequest struct {
	Id int `json:"id"`
}

func DeletePlayer(w http.ResponseWriter, req *http.Request) {
	if b, err := io.ReadAll(req.Body); err != nil {
		handleError(w, err)
		return
	} else {
		var body deletePlayerRequest
		if err := json.Unmarshal(b, &body); err != nil {
			handleError(w, err)
			return
		}
		if err := database.DeletePlayer(body.Id); err != nil {
			handleError(w, err)
			return
		}
		fmt.Fprintf(w, "ok")
	}
}
