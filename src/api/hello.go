package api

import (
	"encoding/json"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) error {
	return json.NewEncoder(w).Encode(&JsonResponse{
		Status: "ok",
		Data:   "hello",
	})
}
