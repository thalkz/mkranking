package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/thalkz/kart/api"
	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/web"
)

var cfg = &config.Config{
	FirstSeasonDate: time.Date(2022, time.September, 17, 0, 0, 0, 0, time.Local),
	SeasonOffset:    3,
	CompetitionDays: 14,
	RestDays:        14,
	MinRacesCount:   5,
	InitialRating:   1000.0,
	Elo: config.ConfigElo{
		K: 32.0,
		D: 400.0,
	},
}

func makeHandler(fn func(*config.Config, http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		err := fn(cfg, w, r)
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
	http.HandleFunc("/history", makeHandler(api.HistoryHandler))
	http.HandleFunc("/player", makeHandler(web.PlayerHandler))
	http.HandleFunc("/results", makeHandler(web.ResultsPageHandler))
	http.HandleFunc("/submit", makeHandler(web.SubmitHandler))
	http.HandleFunc("/new", makeHandler(web.NewPlayerHandler))
	http.HandleFunc("/welcome", makeHandler(web.WelcomePlayerPage))
	http.HandleFunc("/races", makeHandler(web.RacesHandler))
	http.HandleFunc("/stats", makeHandler(web.StatsHandler))
	http.HandleFunc("/champions", makeHandler(web.ChampionsHandler))
	http.HandleFunc("/", makeHandler(web.RankingHandler))

	// Serve static files
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Open database
	var cleanup, err = database.Open(cfg)
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
