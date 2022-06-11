package web

import (
	"fmt"
	"net/http"

	"github.com/thalkz/kart/internal/database"
)

type resultsPage struct {
	Results []result
}

type result struct {
	Rank int
	Name string
	Icon int
	Diff int
}

func ResultsPageHandler(w http.ResponseWriter, r *http.Request) error {
	participantsStr := r.FormValue("participants")
	diffStr := r.FormValue("diff")

	if diffStr == "" || participantsStr == "" {
		return fmt.Errorf("expected diff and participants params")
	}

	playerIds, err := parseIntList(participantsStr)
	if err != nil {
		return fmt.Errorf("failed to parse participants: %w", err)
	}

	diffs, err := parseIntList(diffStr)
	if err != nil {
		return fmt.Errorf("failed to parse diffs: %w", err)
	}

	if len(playerIds) != len(diffs) {
		return fmt.Errorf("expected diff and participants to have the same length")
	}

	players, err := database.GetPlayers(playerIds)
	if err != nil {
		return fmt.Errorf("failed to get players: %w", err)
	}

	results := make([]result, len(players))
	for i := range players {
		results[i] = result{
			Rank: i + 1,
			Name: players[i].Name,
			Icon: players[i].Icon,
			Diff: diffs[i],
		}
	}

	data := resultsPage{
		Results: results,
	}
	renderTemplate(w, "results.html", data)
	return nil
}
