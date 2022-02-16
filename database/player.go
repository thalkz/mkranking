package database

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/thalkz/kart/models"
)

func CreatePlayer(name string, rating float64, icon int) (int, error) {
	row := db.QueryRow("INSERT INTO players (name, rating, icon) values ($1, $2, $3) RETURNING id", name, rating, icon)
	var id int
	err := row.Scan(&id)
	return id, err
}

func DeletePlayer(id int) error {
	_, err := db.Exec("DELETE FROM players where id = $1", id)
	return err
}

func UpdatePlayerName(id int, name string) error {
	_, err := db.Exec("UPDATE players SET name = $1 where id = $2", name, id)
	return err
}

func UpdatePlayerRating(id int, rating float64) error {
	_, err := db.Exec("UPDATE players SET rating = $1 where id = $2", rating, id)
	return err
}

func UpdatePlayerRatings(ids []int, ratings []float64) error {
	if len(ids) != len(ratings) {
		return fmt.Errorf("ids %v and ratings %v should have the same length", ids, ratings)
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for i := range ids {
		_, err := db.Exec("UPDATE players SET rating = $1 where id = $2", ratings[i], ids[i])
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	return err
}

func GetPlayer(id int) (models.Player, error) {
	row := db.QueryRow("SELECT *, RANK () OVER ( ORDER BY rating DESC ) rank FROM players where id = $1", id)
	var player models.Player
	err := row.Scan(&player.Id, &player.Name, &player.Rating, &player.Icon, &player.Rank)
	return player, err
}

func GetPlayers(playerIds []int) ([]models.Player, error) {
	rows, err := db.Query("SELECT *, RANK () OVER ( ORDER BY rating DESC ) rank FROM players WHERE id = ANY($1)", pq.Array(playerIds))
	if err != nil {
		return nil, err
	}

	players := make([]models.Player, 0)
	for rows.Next() {
		var player models.Player
		err = rows.Scan(&player.Id, &player.Name, &player.Rating, &player.Icon, &player.Rank)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	// Order players based on input `playerIds` parameter
	orderedPlayers := make([]models.Player, len(players))
	for i, playerId := range playerIds {
		for _, player := range players {
			if player.Id == playerId {
				orderedPlayers[i] = player
				continue
			}
		}
	}
	return orderedPlayers, err
}

func GetAllPlayers() ([]models.Player, error) {
	rows, err := db.Query("SELECT *, RANK () OVER ( ORDER BY rating DESC ) rank FROM players")
	players := make([]models.Player, 0)
	for rows.Next() {
		var player models.Player
		err = rows.Scan(&player.Id, &player.Name, &player.Rating, &player.Icon, &player.Rank)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, err
}
