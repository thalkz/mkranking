package web

import (
	"fmt"
	"net/http"

	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
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

	for i, race := range races {
		var racePlayers = make([]models.Player, 4)
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

		timeago, err := parseTimeAgo("2006-01-02T15:04:05Z", race.Date)
		if err != nil {
			return fmt.Errorf("failed to parse timeago: %w", err)
		}

		racesWithPlayers[i] = raceWithPlayers{
			Id:      race.Id,
			Date:    timeago,
			Players: racePlayers,
		}
	}

	data := racesPage{
		Races: racesWithPlayers,
	}
	renderTemplate(w, "races.html", data)
	return nil
}
