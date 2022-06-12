package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/internal/database"
)

type resultsPage struct {
	Results []result
}

type result struct {
	Rank int
	Id   int
	Name string
	Icon int
	Diff int
}

func ResultsPageHandler(w http.ResponseWriter, r *http.Request) error {
	raceIdStr := r.FormValue("race_id")
	raceId, err := strconv.Atoi(raceIdStr)
	if err != nil {
		return fmt.Errorf("failed to parse key %v: %w", raceIdStr, err)
	}

	race, err := database.GetRace(raceId)
	if err != nil {
		return fmt.Errorf("failed to get race: %w", err)
	}

	results := make([]result, len(race.Results))
	for i, playerId := range race.Results {
		results[i] = result{
			Rank: i + 1,
			Id:   playerId,
			Name: "??",
			Icon: 0,
			Diff: 0,
		}
	}

	data := resultsPage{
		Results: results,
	}
	renderTemplate(w, "results.html", data)
	return nil
}
