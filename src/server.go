package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/thalkz/kart/src/api"
)

func main() {
	http.HandleFunc("/hello", api.Hello)
	http.HandleFunc("/createPlayer", api.CreatePlayer)
	http.HandleFunc("/submitResults", api.SubmitResults)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	fmt.Println("Listening on port", httpPort)
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
