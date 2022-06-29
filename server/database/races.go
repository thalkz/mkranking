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

func GetRace(id int) (*models.Race, error) {
	rows, err := db.Query(`
	SELECT races.id, 
		races.date, 
		players.id, 
		players.name, 
		players.icon, 
		RANK() OVER (ORDER BY players_races.id), 
		players_races.new_rating, 
		players_races.rating_diff
			FROM races
				JOIN players_races ON players_races.race_id = races.id 
				JOIN players ON players.id = players_races.user_id
			WHERE races.id = $1
			ORDER BY players_races.id DESC`, id)

	race := &models.Race{}
	for rows.Next() {
		var result models.RaceResult
		err = rows.Scan(
			&race.Id,
			&race.Date,
			&result.UserId,
			&result.Name,
			&result.Icon,
			&result.Rank,
			&result.NewRating,
			&result.RatingDiff,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to parse race result: %w", err)
		}
		race.Results[result.Rank-1] = result
	}
	return race, err
}

func GetAllRaces(season int) ([]models.Race, error) {
	rows, err := db.Query(`
		SELECT 
			races.id, 
			races.date, 
			players.id, 
			players.name, 
			players.icon, 
			RANK() OVER (PARTITION BY players_races.race_id ORDER BY players_races.id), 
			players_races.new_rating, 
			players_races.rating_diff
				FROM races
					JOIN players_races ON players_races.race_id = races.id 
					JOIN players ON players.id = players_races.user_id
				WHERE races.season = $1
				ORDER BY players_races.id DESC`, season)
	races := make([]models.Race, 0)
	previousRaceId := 0
	raceId := previousRaceId
	raceDate := ""
	var race *models.Race

	for rows.Next() {
		var result models.RaceResult
		err = rows.Scan(
			&raceId,
			&raceDate,
			&result.UserId,
			&result.Name,
			&result.Icon,
			&result.Rank,
			&result.NewRating,
			&result.RatingDiff,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to parse race result: %w", err)
		}

		// Check if row is a new race
		if raceId != previousRaceId {
			if race != nil {
				races = append(races, *race)
			}
			race = &models.Race{}
			race.Id = raceId
			race.Date = raceDate
			previousRaceId = raceId
		}
		race.Results[result.Rank-1] = result
	}

	// Append last race
	if race != nil {
		races = append(races, *race)
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
	rows, err := db.Query(`
		SELECT selected_races.race_id, 
			selected_races.date, 
			players.id, players.name, 
			players.icon, RANK() OVER (PARTITION BY players_races.race_id ORDER BY players_races.id), 
			players_races.new_rating, 
			players_races.rating_diff
				FROM (SELECT * FROM races 
							JOIN players_races ON players_races.race_id = races.id 
							WHERE user_id = $1
							ORDER BY date) as selected_races
					JOIN players_races ON players_races.race_id = selected_races.race_id 
					JOIN players ON players.id = players_races.user_id
				ORDER BY players_races.id DESC`, userId)

	races := make([]models.Race, 0)
	var previousRaceId int
	var raceId int
	var raceDate string
	var race *models.Race

	for rows.Next() {
		var result models.RaceResult
		err = rows.Scan(
			&raceId,
			&raceDate,
			&result.UserId,
			&result.Name,
			&result.Icon,
			&result.Rank,
			&result.NewRating,
			&result.RatingDiff,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to parse race result: %w", err)
		}

		// Check if row is a new race
		if raceId != previousRaceId {
			if race != nil {
				races = append(races, *race)
			}
			race = &models.Race{}
			race.Id = raceId
			race.Date = raceDate
			previousRaceId = raceId
		}
		race.Results[result.Rank-1] = result
	}

	// Append last race
	if race != nil {
		races = append(races, *race)
	}
	return races, err
}
