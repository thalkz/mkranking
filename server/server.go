package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"

	_ "github.com/lib/pq"
	"github.com/thalkz/kart/api"
	"github.com/thalkz/kart/database"
)

func appHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		err := fn(w, r)
		end := time.Now()
		var statusCode = http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
			bytes, _ := json.Marshal(&api.JsonResponse{
				Status: "error",
				Error:  err.Error(),
			})
			log.Println("Error:", r.URL, err)
			http.Error(w, string(bytes), statusCode)
		}

		if r.Method != "OPTIONS" {
			duration := end.UnixMilli() - start.UnixMilli()
			log.Printf("%v %v %v (%vms)\n", r.RemoteAddr, r.URL, http.StatusText(statusCode), duration)
		}
	}
}

func main() {
	router := http.NewServeMux()

	// Setup routes
	router.HandleFunc("/hello", appHandler(api.Hello))
	router.HandleFunc("/getPlayer", appHandler(api.GetPlayer))
	router.HandleFunc("/createPlayer", appHandler(api.CreatePlayer))
	router.HandleFunc("/updatePlayer", appHandler(api.UpdatePlayer))
	router.HandleFunc("/deletePlayer", appHandler(api.DeletePlayer))
	router.HandleFunc("/submitResults", appHandler(api.SubmitResults))
	router.HandleFunc("/getAllPlayers", appHandler(api.GetAllPlayers))
	router.HandleFunc("/getAllRaces", appHandler(api.GetAllRaces))
	router.HandleFunc("/getPlayerRaces", appHandler(api.GetPlayerRaces))
	router.HandleFunc("/resetAllRatings", appHandler(api.ResetAllRatings))
	router.HandleFunc("/getHistory", appHandler(api.GetHistory))

	// Open database
	var cleanup, err = database.Open()
	if err != nil {
		log.Fatalln("failed to open database:", err)
	}
	log.Println("Database is opened")
	defer cleanup()

	// Add CORS middleware
	c := cors.AllowAll()
	handler := c.Handler(router)

	// Get port
	httpPort := os.Getenv("SERVER_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}

	// Start server
	log.Println("Listening on port", httpPort)
	err = http.ListenAndServe(":"+httpPort, handler)
	if err != nil {
		log.Fatalln("listen and server failed:", err)
	}
}
