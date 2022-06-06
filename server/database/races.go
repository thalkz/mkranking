package database

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/thalkz/kart/models"
)

// ranking is a list of playerIds, sorted by race position
func CreateRace(ranking []int, oldRatings, newRatings []float64) error {
	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}

	// Insert race
	now := time.Now().Format("2006-01-02 15:04:05")
	row := tx.QueryRow("INSERT INTO races (ranking, date) VALUES ($1, $2) RETURNING id", pq.Array(ranking), now)
	var raceId int
	if err = row.Scan(&raceId); err != nil {
		return fmt.Errorf("failed to insert into races: %w", err)
	}

	// Insert players_races
	for i, userId := range ranking {
		_, err := tx.Exec(
			`INSERT INTO players_races (user_id, race_id, old_rating, new_rating, rating_diff) 
				VALUES ($1, $2, $3, $4, $5) 
				RETURNING rating_diff`,
			userId,
			raceId,
			oldRatings[i],
			newRatings[i],
			newRatings[i]-oldRatings[i],
		)
		if err != nil {
			return fmt.Errorf("failed to insert %v into players_races: %w", userId, err)
		}
	}

	// Update players with current rating
	for i, userId := range ranking {
		_, err := db.Exec("UPDATE players SET rating = $1 WHERE id = $2", newRatings[i], userId)
		if err != nil {
			return fmt.Errorf("failed to update player %v: %w", userId, err)
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx: %w", err)
	}
	return nil
}

func GetRace(id int) (models.Race, error) {
	row := db.QueryRow("SELECT * FROM races WHERE id = $1", id)
	var race models.Race
	var ranking pq.Int64Array
	err := row.Scan(&race.Id, &ranking, &race.Date)
	for i := range ranking {
		race.Results = append(race.Results, (int)(ranking[i]))
	}
	return race, err
}

func GetAllRaces() ([]models.Race, error) {
	rows, err := db.Query("SELECT * FROM races")
	races := make([]models.Race, 0)
	var ranking pq.Int64Array
	for rows.Next() {
		var race models.Race
		err = rows.Scan(&race.Id, &ranking, &race.Date)
		if err != nil {
			return nil, err
		}
		for i := range ranking {
			race.Results = append(race.Results, (int)(ranking[i]))
		}
		races = append(races, race)
	}
	return races, err
}
