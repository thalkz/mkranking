package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/thalkz/kart/src/api"
	"github.com/thalkz/kart/src/database"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		bytes, _ := json.Marshal(&api.JsonResponse{
			Status: "error",
			Error:  err,
		})
		http.Error(w, string(bytes), 500)
	}
}

func main() {
	http.Handle("/hello", appHandler(api.Hello))
	http.Handle("/getPlayer", appHandler(api.GetPlayer))
	http.Handle("/createPlayer", appHandler(api.CreatePlayer))
	http.Handle("/updatePlayer", appHandler(api.UpdatePlayer))
	http.Handle("/deletePlayer", appHandler(api.DeletePlayer))
	http.Handle("/submitResults", appHandler(api.SubmitResults))
	http.Handle("/getAllPlayers", appHandler(api.GetAllPlayers))
	http.Handle("/ResetAllRatings", appHandler(api.ResetAllRatings))

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	err := database.Open()
	if err != nil {
		panic(err)
	}
	if err = database.CreatePlayersTable(); err != nil {
		panic(err)
	}
	if err = database.CreateRacesTable(); err != nil {
		panic(err)
	}
	if err = database.CreatePlayersRacesTable(); err != nil {
		panic(err)
	}

	defer database.Close()

	fmt.Println("Listening on port", httpPort)
	err = http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		panic(err)
	}
}
