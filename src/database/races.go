package database

import (
	"fmt"
	"strings"

	"github.com/thalkz/kart/src/models"
)

func CreateRace(ranking []int) error {
	rankingArray := strings.Trim(strings.Join(strings.Split(fmt.Sprint(ranking), " "), ", "), "[]")
	statement := fmt.Sprintf(`INSERT INTO races (ranking) values (ARRAY[%v]) RETURNING id;`, rankingArray)
	fmt.Println(statement)
	row := db.QueryRow(statement)
	var raceId int
	err := row.Scan(&raceId)
	if err != nil {
		return err
	}
	for _, userId := range ranking {
		linkStatement := fmt.Sprintf(`INSERT INTO players_races (user_id, race_id) values (%v, %v) RETURNING id;`, userId, raceId)
		_, linkErr := db.Exec(linkStatement)
		if linkErr != nil {
			return linkErr
		}
	}
	return err
}

func GetRace(id int) (models.Race, error) {
	statement := fmt.Sprintf(`SELECT * FROM races where id = %v`, id)
	row := db.QueryRow(statement)
	var race models.Race
	err := row.Scan(&race.Id, &race.Ranking)
	return race, err
}
