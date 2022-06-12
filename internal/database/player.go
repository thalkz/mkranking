package database

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/thalkz/kart/internal/models"
)

func CreatePlayer(name string, rating float64, icon int) (int, error) {
	row := db.QueryRow("INSERT INTO players (name, rating, icon) VALUES ($1, $2, $3) RETURNING id", name, rating, icon)
	var id int
	err := row.Scan(&id)
	return id, err
}

func DeletePlayer(id int) error {
	_, err := db.Exec("DELETE FROM players WHERE id = $1", id)
	return err
}

func UpdatePlayerName(id int, name string) error {
	_, err := db.Exec("UPDATE players SET name = $1 WHERE id = $2", name, id)
	return err
}

func UpdatePlayerRating(id int, rating float64) error {
	_, err := db.Exec("UPDATE players SET rating = $1 WHERE id = $2", rating, id)
	return err
}

func ResetAllRatings(rating float64) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}
	defer tx.Rollback()

	_, err = db.Exec("UPDATE players SET rating = $1", rating)
	if err != nil {
		return fmt.Errorf("failed to update players: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit tx: %w", err)
	}
	return nil
}

func GetPlayer(id int) (models.Player, error) {
	row := db.QueryRow("SELECT id, name, rating, icon, races_count, RANK() OVER (ORDER BY rating DESC) rank FROM players WHERE id = $1", id)
	var player models.Player
	err := row.Scan(&player.Id, &player.Name, &player.Rating, &player.Icon, &player.RacesCount, &player.Rank)
	return player, err
}

func GetPlayers(playerIds []int) ([]models.Player, error) {
	rows, err := db.Query("SELECT id, name, rating, icon, races_count, RANK() OVER (ORDER BY rating DESC) rank FROM players WHERE id = ANY($1)", pq.Array(playerIds))
	if err != nil {
		return nil, err
	}

	players := make([]models.Player, 0)
	for rows.Next() {
		var player models.Player
		err = rows.Scan(&player.Id, &player.Name, &player.Rating, &player.Icon, &player.RacesCount, &player.Rank)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	if len(playerIds) != len(players) {
		return nil, fmt.Errorf("returned players does not have same length: expected %v, got %v", len(playerIds), len(players))
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
	rows, err := db.Query("SELECT id, name, rating, icon, races_count, RANK() OVER (ORDER BY rating DESC) rank FROM players")
	players := make([]models.Player, 0)
	for rows.Next() {
		var player models.Player
		err = rows.Scan(&player.Id, &player.Name, &player.Rating, &player.Icon, &player.RacesCount, &player.Rank)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, err
}

func GetRankedPlayers(minRaces int) ([]models.Player, error) {
	rows, err := db.Query("SELECT id, name, rating, icon, races_count, RANK() OVER (ORDER BY rating DESC) rank FROM players WHERE races_count >= $1", minRaces)
	players := make([]models.Player, 0)
	for rows.Next() {
		var player models.Player
		err = rows.Scan(&player.Id, &player.Name, &player.Rating, &player.Icon, &player.RacesCount, &player.Rank)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, err
}

func GetUnrankedPlayers(minRaces int) ([]models.Player, error) {
	rows, err := db.Query("SELECT id, name, rating, icon, races_count, RANK() OVER (ORDER BY rating DESC) rank FROM players WHERE races_count < $1", minRaces)
	players := make([]models.Player, 0)
	for rows.Next() {
		var player models.Player
		err = rows.Scan(&player.Id, &player.Name, &player.Rating, &player.Icon, &player.RacesCount, &player.Rank)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, err
}
