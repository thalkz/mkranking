package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/web"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		err := fn(w, r)
		end := time.Now()

		var statusCode = http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
			log.Println("Error:", r.URL, err)
			http.Error(w, err.Error(), statusCode)
		}

		duration := end.UnixMilli() - start.UnixMilli()
		log.Printf("%v %v %v (%vms)\n", r.RemoteAddr, r.URL, http.StatusText(statusCode), duration)
	}
}

func main() {
	// Serve routes
	http.HandleFunc("/player", makeHandler(web.PlayerHandler))
	http.HandleFunc("/results", makeHandler(web.ResultsPageHandler))
	http.HandleFunc("/submit", makeHandler(web.SubmitHandler))
	http.HandleFunc("/new", makeHandler(web.NewPlayerHandler))
	http.HandleFunc("/welcome", makeHandler(web.WelcomePlayerPage))
	http.HandleFunc("/races", makeHandler(web.RacesHandler))
	http.HandleFunc("/stats", makeHandler(web.StatsHandler))
	http.HandleFunc("/", makeHandler(web.RankingHandler))

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
