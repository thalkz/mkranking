package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/web"
)

// func appHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		err := fn(w, r)
// 		end := time.Now()
// 		var statusCode = http.StatusOK
// 		if err != nil {
// 			statusCode = http.StatusInternalServerError
// 			bytes, _ := json.Marshal(&api.JsonResponse{
// 				Status: "error",
// 				Error:  err.Error(),
// 			})
// 			log.Println("Error:", r.URL, err)
// 			http.Error(w, string(bytes), statusCode)
// 		}

// 		if r.Method != "OPTIONS" {
// 			duration := end.UnixMilli() - start.UnixMilli()
// 			log.Printf("%v %v %v (%vms)\n", r.RemoteAddr, r.URL, http.StatusText(statusCode), duration)
// 		}
// 	}
// }

func main() {
	// Serve routes
	http.HandleFunc("/submit", web.SubmitResultsHandler)
	http.HandleFunc("/new", web.NewPlayerHandler)
	http.HandleFunc("/player", web.PlayerHandler)
	http.HandleFunc("/races", web.RacesHandler)
	http.HandleFunc("/stats", web.StatsHandler)
	http.HandleFunc("/", web.RankingHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Open database
	var cleanup, err = database.Open()
	if err != nil {
		log.Fatalln("failed to open database:", err)
	}
	defer cleanup()

	// Get port
	httpPort := os.Getenv("SERVER_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}

	// Start server
	log.Println("Listening on port", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
