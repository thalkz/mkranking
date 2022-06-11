package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/thalkz/kart/internal/database"
	"github.com/thalkz/kart/internal/elo"
	"github.com/thalkz/kart/internal/models"
)

type participantsPage struct {
	Players []models.Player
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) error {
	participants := r.FormValue("participants")
	confirm := r.FormValue("confirm")

	if confirm == "true" {
		return createRaceHandler(w, r, participants)
	} else if participants != "" {
		return orderParticipantsHandler(w, r, participants)
	} else {
		return selectParticipantsHandler(w, r)
	}
}

func selectParticipantsHandler(w http.ResponseWriter, r *http.Request) error {
	players, err := database.GetAllPlayers()
	if err != nil {
		return fmt.Errorf("failed to get all players: %w", err)
	}

	data := participantsPage{
		Players: players,
	}
	renderTemplate(w, "submit_select.html", data)
	return nil
}

func orderParticipantsHandler(w http.ResponseWriter, r *http.Request, participants string) error {
	ids, err := parseIntList(participants)
	if err != nil {
		return fmt.Errorf("failed to parse participants: %w", err)
	}

	players, err := database.GetPlayers(ids)
	if err != nil {
		return fmt.Errorf("failed to get players: %w", err)
	}

	data := participantsPage{
		Players: players,
	}
	renderTemplate(w, "submit_order.html", data)
	return nil
}

func createRaceHandler(w http.ResponseWriter, r *http.Request, participants string) error {
	ids, err := parseIntList(participants)
	if err != nil {
		return fmt.Errorf("failed to parse participants: %w", err)
	}

	if len(ids) < 2 {
		return fmt.Errorf("cannot submit results with less than two players. recieved %v", ids)
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

	if err := database.CreateRace(ids, oldRatings, newRatings); err != nil {
		return fmt.Errorf("failed creating race: %w", err)
	}
	log.Printf("Created race with participants: %v. Updated ratings to %v\n", ids, newRatings)

	var ratingDiff = make([]string, len(newRatings))
	for i := range players {
		ratingDiff[i] = strconv.Itoa(int(newRatings[i] - oldRatings[i]))
	}
	diff := strings.Join(ratingDiff, ",")

	http.Redirect(w, r, fmt.Sprintf("/results?participants=%v&diff=%v", participants, diff), http.StatusFound)

	return nil
}
