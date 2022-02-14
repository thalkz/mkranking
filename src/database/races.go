package database

import (
	"github.com/lib/pq"
	"github.com/thalkz/kart/src/models"
)

func CreateRace(ranking []int) error {
	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	row := tx.QueryRow("INSERT INTO races (ranking) values ($1) RETURNING id", pq.Array(ranking))
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
	err := row.Scan(&race.Id, &race.Ranking)
	return race, err
}
