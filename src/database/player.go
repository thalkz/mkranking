package database

import (
	"fmt"

	"github.com/thalkz/kart/src/models"
)

func CreatePlayer(name string, rating float64) (int, error) {
	statement := fmt.Sprintf(`INSERT INTO "Players"(name, rating) values('%s', %v) RETURNING id;`, name, rating)
	row := db.QueryRow(statement)
	var id int
	err := row.Scan(&id)
	return id, err
}

func DeletePlayer(id int) error {
	statement := fmt.Sprintf(`DELETE FROM "Players" where id = %v`, id)
	_, err := db.Exec(statement)
	return err
}

func UpdatePlayerName(id int, name string) error {
	statement := fmt.Sprintf(`UPDATE "Players" SET name = '%s' where id = %v;`, name, id)
	_, err := db.Exec(statement)
	return err
}

func UpdatePlayerRating(id int, rating float64) error {
	statement := fmt.Sprintf(`UPDATE "Players" SET rating = %v where id = %v;`, rating, id)
	_, err := db.Exec(statement)
	return err
}

func GetPlayer(id int) (models.Player, error) {
	statement := fmt.Sprintf(`SELECT * FROM "Players" where id = %v`, id)
	row := db.QueryRow(statement)
	var player models.Player
	err := row.Scan(&player.Id, &player.Name, &player.Rating)
	return player, err
}

func GetPlayers() ([]models.Player, error) {
	statement := fmt.Sprintf(`SELECT * FROM "Players"`)
	rows, err := db.Query(statement)
	players := make([]models.Player, 0)
	for next := rows.Next(); next; next = rows.Next() {
		var player models.Player
		err = rows.Scan(&player.Id, &player.Name, &player.Rating)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, err
}
