package database

import (
	"database/sql"
	"fmt"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/models"
)

func GetAllPlayersEvents(season int) (map[int][]models.HistoryEvent, error) {
	rows, err := db.Query(
		`SELECT players.id AS player_id, races.id AS race_id, date, new_rating
			FROM (SELECT * FROM races WHERE season = $1) as races
				CROSS JOIN (SELECT * FROM players WHERE season = $1) as players
				LEFT JOIN players_races ON players.id = players_races.user_id AND races.id = players_races.race_id
				ORDER BY date`, season)
	if err != nil {
		return nil, fmt.Errorf("failed to query all history: %w", err)
	}

	events := make(map[int][]models.HistoryEvent)
	lastValidRating := config.InitialRating
	var previousPlayerId int
	for rows.Next() {
		var raceId int
		var date string
		var newRating sql.NullFloat64 // New rating is null when player did not participate in race
		var playerId int
		err = rows.Scan(&playerId, &raceId, &date, &newRating)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if newRating.Valid {
			lastValidRating = newRating.Float64
		} else if previousPlayerId != playerId {
			lastValidRating = config.InitialRating
		}

		events[playerId] = append(events[playerId], models.HistoryEvent{
			RaceId: raceId,
			Rating: lastValidRating,
		})

		if previousPlayerId != playerId {
			previousPlayerId = playerId
		}
	}

	// TODO Verify that all event arrays have the same length

	return events, nil
}

func GetPlayerHistory(userId int) ([]models.HistoryEvent, error) {
	rows, err := db.Query(
		`SELECT races.id, date, new_rating 
			FROM races LEFT JOIN players_races 
				ON user_id = $1 AND players_races.race_id = races.id 
			ORDER BY date`, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query history: %w", err)
	}

	history := make([]models.HistoryEvent, 0)
	lastValidRating := config.InitialRating
	for rows.Next() {
		var raceId int
		var date string
		var newRating sql.NullFloat64 // New rating is null when player did not participate in race
		err = rows.Scan(&raceId, &date, &newRating)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if newRating.Valid {
			lastValidRating = newRating.Float64
		}

		history = append(history, models.HistoryEvent{
			RaceId: raceId,
			Rating: lastValidRating,
		})
	}

	return history, nil
}
