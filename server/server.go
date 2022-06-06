package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/cors"

	_ "github.com/lib/pq"
	"github.com/thalkz/kart/api"
	"github.com/thalkz/kart/database"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	err := fn(w, r)
	end := time.Now()
	logRequest(r, start, end)
	if err != nil {
		bytes, _ := json.Marshal(&api.JsonResponse{
			Status: "error",
			Error:  err.Error(),
		})
		fmt.Println("Error: ", r.URL, err)
		http.Error(w, string(bytes), 500)
	}
}

func logRequest(r *http.Request, start, end time.Time) {
	ipAddr := strings.Split(r.RemoteAddr, ":")[0]
	duration := end.UnixMilli() - start.UnixMilli()
	fmt.Printf("[%v] %v %v %v (%vms)\n", start.Format("2006-01-02 15:04:05"), ipAddr, r.Method, r.URL, duration)
}

func main() {
	router := http.NewServeMux()

	// Setup routes
	router.Handle("/hello", appHandler(api.Hello))
	router.Handle("/getPlayer", appHandler(api.GetPlayer))
	router.Handle("/createPlayer", appHandler(api.CreatePlayer))
	router.Handle("/updatePlayer", appHandler(api.UpdatePlayer))
	router.Handle("/deletePlayer", appHandler(api.DeletePlayer))
	router.Handle("/submitResults", appHandler(api.SubmitResults))
	router.Handle("/getAllPlayers", appHandler(api.GetAllPlayers))
	router.Handle("/getAllRaces", appHandler(api.GetAllRaces))
	router.Handle("/resetAllRatings", appHandler(api.ResetAllRatings))
	router.Handle("/getRatingsHistory", appHandler(api.GetRatingsHistory))

	// Open database
	var cleanup, err = database.Open()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database is opened")
	defer cleanup()

	// Add CORS middleware
	c := cors.Default()
	handler := c.Handler(router)

	// Get port
	httpPort := os.Getenv("SERVER_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}

	// Start server
	fmt.Println("Listening on port", httpPort)
	err = http.ListenAndServe(":"+httpPort, handler)
	if err != nil {
		panic(err)
	}
}
