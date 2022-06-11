package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/thalkz/kart/internal/database"
	"github.com/thalkz/kart/internal/models"
)

type racesPage struct {
	Races []raceWithPlayers
}

type raceWithPlayers struct {
	Id      int
	Players []models.Player
	Date    string
}

func RacesHandler(w http.ResponseWriter, r *http.Request) error {
	races, err := database.GetAllRaces()
	if err != nil {
		return fmt.Errorf("failed to get all races: %w", err)
	}

	players, err := database.GetAllPlayers()
	if err != nil {
		return fmt.Errorf("failed to get all players: %w", err)
	}

	var playersMap = make(map[int]models.Player)
	for _, player := range players {
		playersMap[player.Id] = player
	}

	racesWithPlayers := make([]raceWithPlayers, len(races))

	now := time.Now().Unix()

	for i, race := range races {
		var racePlayers = make([]models.Player, len(race.Results))
		for i, playerId := range race.Results {
			value, ok := playersMap[playerId]
			if ok {
				racePlayers[i] = value
			} else {
				racePlayers[i] = models.Player{
					Name: "[deleted]",
				}
			}
		}

		date, err := time.Parse("2006-01-02T15:04:05Z", race.Date)
		if err != nil {
			return fmt.Errorf("failed to parse race date: %w", err)
		}

		second := now - date.Unix()
		hours := int(second / 3600)
		days := int(hours / 24)

		var dateStr string
		if days > 0 {
			dateStr = fmt.Sprintf("%v jours", days)
		} else if hours > 0 {
			dateStr = fmt.Sprintf("%v heures", days)
		} else {
			dateStr = "à l'instant"
		}

		racesWithPlayers[i] = raceWithPlayers{
			Id:      race.Id,
			Date:    dateStr,
			Players: racePlayers,
		}
	}

	data := racesPage{
		Races: racesWithPlayers,
	}
	renderTemplate(w, "races.html", data)
	return nil
}
