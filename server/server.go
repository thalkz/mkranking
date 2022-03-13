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
	err := fn(w, r)
	logRequest(r)
	if err != nil {
		bytes, _ := json.Marshal(&api.JsonResponse{
			Status: "error",
			Error:  err.Error(),
		})
		fmt.Println("Error: ", r.URL, err)
		http.Error(w, string(bytes), 500)
	}
}

func logRequest(r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	ipAddr := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Printf("[%v] %v %v %v\n", now, ipAddr, r.Method, r.URL)
}

func main() {
	mux := http.NewServeMux()

	// Add routes
	mux.Handle("/hello", appHandler(api.Hello))
	mux.Handle("/getPlayer", appHandler(api.GetPlayer))
	mux.Handle("/createPlayer", appHandler(api.CreatePlayer))
	mux.Handle("/updatePlayer", appHandler(api.UpdatePlayer))
	mux.Handle("/deletePlayer", appHandler(api.DeletePlayer))
	mux.Handle("/submitResults", appHandler(api.SubmitResults))
	mux.Handle("/getAllPlayers", appHandler(api.GetAllPlayers))
	mux.Handle("/getAllRaces", appHandler(api.GetAllRaces))
	mux.Handle("/resetAllRatings", appHandler(api.ResetAllRatings))
	mux.Handle("/getRatingsHistory", appHandler(api.GetRatingsHistory))

	// Open database
	if err := database.Open(); err != nil {
		panic(err)
	}
	defer database.Close()

	// Enable CORS
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://kart.thalkz.com"},
	// 	AllowCredentials: true,
	// 	// Enable Debugging for testing, consider disabling in production
	// 	Debug: true,
	// })
	c := cors.AllowAll()
	handler := c.Handler(mux)

	// Get port
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "443"
	}

	// Start server
	fmt.Println("Listening on port", httpPort)
	err := http.ListenAndServe(":"+httpPort, handler)
	if err != nil {
		panic(err)
	}
}
