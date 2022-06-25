package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/elo"
	"github.com/thalkz/kart/models"
)

type participantsPage struct {
	Players []models.Player
}

const NO_PLAYER_ID = -1

func SubmitHandler(w http.ResponseWriter, r *http.Request) error {
	first := r.FormValue("first")

	if first != "" {
		return createRaceHandler(w, r)
	} else {
		return selectParticipantsHandler(w, r)
	}
}

func selectParticipantsHandler(w http.ResponseWriter, r *http.Request) error {
	players, err := database.GetAllPlayers(config.Season)
	if err != nil {
		return fmt.Errorf("failed to get all players: %w", err)
	}

	data := participantsPage{
		Players: players,
	}
	renderTemplate(w, "submit.html", data)
	return nil
}

func createRaceHandler(w http.ResponseWriter, r *http.Request) error {
	ids, err := parseParticipantsKeys(r)
	if err != nil {
		return fmt.Errorf("failed to parse participants: %w", err)
	}

	players, err := database.GetPlayers(ids)
	if err != nil {
		return fmt.Errorf("failed getting players: %w", err)
	}

	oldRatings := make([]float64, len(players))
	for i := range players {
		oldRatings[i] = players[i].Rating
	}

	newRatings := elo.ComputeRatings(oldRatings)

	raceId, err := database.CreateRace(ids, oldRatings, newRatings, config.Season)
	if err != nil {
		return fmt.Errorf("failed creating race: %w", err)
	}
	log.Printf("Created race with participants %v\n", ids)
	log.Printf("Updated ratings to %v\n", newRatings)

	endpoint := fmt.Sprintf("/results?race_id=%v&show_ok_button=true", raceId)
	http.Redirect(w, r, endpoint, http.StatusFound)

	return nil
}

func parseParticipantsKeys(r *http.Request) ([]int, error) {
	ids := make([]int, 0)
	ids, err := appendParticipantKey(ids, r, "first", true)
	if err != nil {
		return nil, fmt.Errorf("failed to append first key: %w", err)
	}

	ids, err = appendParticipantKey(ids, r, "second", true)
	if err != nil {
		return nil, fmt.Errorf("failed to append second key: %w", err)
	}

	ids, err = appendParticipantKey(ids, r, "third", false)
	if err != nil {
		return nil, fmt.Errorf("failed to append third key: %w", err)
	}

	ids, err = appendParticipantKey(ids, r, "fourth", false)
	if err != nil {
		return nil, fmt.Errorf("failed to append fourth key: %w", err)
	}
	return ids, nil
}

func appendParticipantKey(ids []int, r *http.Request, key string, isRequired bool) ([]int, error) {
	valueStr := r.FormValue(key)
	if valueStr == "" {
		if isRequired {
			return nil, fmt.Errorf("required key is empty: %v", key)
		}
		return ids, nil
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse key %v: %w", key, err)
	}
	for _, id := range ids {
		if id == value {
			return nil, fmt.Errorf("participant %v already exists in %v", value, ids)
		}
	}
	return append(ids, value), nil
}
