package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thalkz/kart/api"
)

func ErrorHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			bytes, _ := json.Marshal(&api.JsonResponse{
				Status: "error",
				Error:  err.Error(),
			})
			fmt.Println("Error: ", r.URL, err)
			http.Error(w, string(bytes), 500)
		}
	}
}
