package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/thalkz/kart/src/api"
	"github.com/thalkz/kart/src/database"
)

var Database *sql.DB

func main() {
	http.HandleFunc("/hello", api.Hello)
	http.HandleFunc("/getPlayer", api.GetPlayer)
	http.HandleFunc("/createPlayer", api.CreatePlayer)
	http.HandleFunc("/updatePlayer", api.UpdatePlayer)
	http.HandleFunc("/deletePlayer", api.DeletePlayer)
	http.HandleFunc("/submitResults", api.SubmitResults)
	http.HandleFunc("/getPlayers", api.GetPlayers)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	err := database.Open()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	fmt.Println("Listening on port", httpPort)
	err = http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		panic(err)
	}
}
