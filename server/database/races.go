package database

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/thalkz/kart/models"
)

// ranking is a list of playerIds, sorted by race position
func CreateRace(ranking []int, oldRatings, newRatings []float64, season int) (int, error) {
	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return 0, fmt.Errorf("failed to begin tx: %w", err)
	}

	// Insert race
	now := time.Now().Format("2006-01-02 15:04:05")
	row := tx.QueryRow("INSERT INTO races (ranking, date, season) VALUES ($1, $2, $3) RETURNING id", pq.Array(ranking), now, season)
	var raceId int
	if err = row.Scan(&raceId); err != nil {
		return 0, fmt.Errorf("failed to insert into races: %w", err)
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
			return 0, fmt.Errorf("failed to insert %v into players_races: %w", userId, err)
		}
	}

	// Update players with current rating
	for i, userId := range ranking {
		_, err := db.Exec("UPDATE players SET rating = $1, races_count = races_count + 1 WHERE id = $2", newRatings[i], userId)
		if err != nil {
			return 0, fmt.Errorf("failed to update player %v: %w", userId, err)
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit tx: %w", err)
	}
	return raceId, nil
}

func GetRace(id int) (models.Race, error) {
	row := db.QueryRow(`SELECT id, ranking, date FROM races WHERE races.id = $1`, id)
	var race models.Race
	var ranking pq.Int64Array
	err := row.Scan(&race.Id, &ranking, &race.Date)
	for i := range ranking {
		race.Results = append(race.Results, (int)(ranking[i]))
	}
	return race, err
}

func GetRaceDetails(id int) (*models.RaceDetails, error) {
	race, err := GetRace(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get race: %w", err)
	}
	rows, err := db.Query(`SELECT players_races.user_id, players.name, players.icon, players_races.new_rating, players_races.rating_diff, RANK() OVER (ORDER BY players_races.id)
		FROM races 
			JOIN players_races ON players_races.race_id = races.id 
			JOIN players ON players.id = players_races.user_id
		WHERE races.id = $1
		ORDER BY players_races.id`, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query race details: %w", err)
	}
	results := make([]models.RaceDetailsResult, 0)
	for rows.Next() {
		var result models.RaceDetailsResult
		err = rows.Scan(&result.UserId, &result.Name, &result.Icon, &result.NewRating, &result.RatingDiff, &result.RaceRank)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, result)
	}
	return &models.RaceDetails{
		Race:    race,
		Results: results,
	}, nil
}

func GetAllRaces(season int) ([]models.Race, error) {
	rows, err := db.Query("SELECT id, ranking, date FROM races WHERE season = $1 ORDER BY date DESC", season)
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

func GetPlayerTotalRaceCount(userId int) (int, error) {
	row := db.QueryRow("SELECT COUNT(*) FROM players_races WHERE user_id = $1", userId)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to count total races for %v: %w", userId, err)
	}
	return count, nil
}

func GetPlayerRaces(userId int) ([]models.Race, error) {
	rows, err := db.Query(`SELECT races.id, races.ranking, races.date
		FROM races JOIN players_races ON players_races.race_id = races.id 
		WHERE user_id = $1
		ORDER BY date`, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query races for player %v: %w", userId, err)
	}

	races := make([]models.Race, 0)
	var ranking pq.Int64Array
	for rows.Next() {
		var race models.Race
		err = rows.Scan(&race.Id, &ranking, &race.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan race row: %w", err)
		}
		for i := range ranking {
			race.Results = append(race.Results, (int)(ranking[i]))
		}
		races = append(races, race)
	}
	return races, nil
}
