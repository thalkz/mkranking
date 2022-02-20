package database

import (
	"time"

	"github.com/lib/pq"
	"github.com/thalkz/kart/models"
)

func CreateRace(ranking []int) error {
	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	row := tx.QueryRow("INSERT INTO races (ranking, date) values ($1, $2) RETURNING id", pq.Array(ranking), now)
	var raceId int

	if err = row.Scan(&raceId); err != nil {
		return err
	}
	for _, userId := range ranking {
		_, err := tx.Exec("INSERT INTO players_races (user_id, race_id) values ($1, $2) RETURNING id", userId, raceId)
		if err != nil {
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return err
}

func GetRace(id int) (models.Race, error) {
	row := db.QueryRow("SELECT * FROM races where id = $1", id)
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
