package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	_ "github.com/lib/pq"
	"github.com/thalkz/kart/api"
	"github.com/thalkz/kart/database"
	mw "github.com/thalkz/kart/middlewares"
)

func main() {
	router := http.NewServeMux()

	// Setup routes
	router.HandleFunc("/hello", mw.ErrorHandler(api.Hello))
	router.HandleFunc("/getPlayer", mw.ErrorHandler(api.GetPlayer))
	router.HandleFunc("/createPlayer", mw.ErrorHandler(api.CreatePlayer))
	router.HandleFunc("/updatePlayer", mw.ErrorHandler(api.UpdatePlayer))
	router.HandleFunc("/deletePlayer", mw.ErrorHandler(api.DeletePlayer))
	router.HandleFunc("/submitResults", mw.ErrorHandler(api.SubmitResults))
	router.HandleFunc("/getAllPlayers", mw.ErrorHandler(api.GetAllPlayers))
	router.HandleFunc("/getAllRaces", mw.ErrorHandler(api.GetAllRaces))
	router.HandleFunc("/getPlayerRaces", mw.ErrorHandler(api.GetPlayerRaces))
	router.HandleFunc("/resetAllRatings", mw.ErrorHandler(api.ResetAllRatings))
	router.HandleFunc("/getRatingsHistory", mw.ErrorHandler(api.GetRatingsHistory))

	// Open database
	var cleanup, err = database.Open()
	if err != nil {
		panic(err)
	}
	log.Println("Database is opened")
	defer cleanup()

	// Add CORS middleware
	c := cors.Default()
	handler := c.Handler(router)

	// Add logger middleware
	handler = mw.LoggerHandler(handler)

	// Get port
	httpPort := os.Getenv("SERVER_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}

	// Start server
	log.Println("Listening on port", httpPort)
	err = http.ListenAndServe(":"+httpPort, handler)
	if err != nil {
		panic(err)
	}
}
