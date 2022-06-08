package database

import (
	"database/sql"
	"fmt"

	"github.com/thalkz/kart/internal/models"
)

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
	lastValidRating := 1000.0 // TODO Use initial rating config
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
			RaceId:    raceId,
			Date:      date,
			NewRating: lastValidRating,
		})
	}

	return history, nil
}
