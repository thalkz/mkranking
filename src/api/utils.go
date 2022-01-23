package api

import (
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	fmt.Printf("Error: %v\n", err)
	fmt.Fprintf(w, "Error: %v", err)
}
