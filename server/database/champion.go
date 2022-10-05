package database

import "github.com/thalkz/kart/models"

func GetChampions(season int) ([]models.Champion, error) {
	rows, err := db.Query(`
	WITH ranked_players AS 
		( SELECT season, name, icon, RANK() OVER (PARTITION BY season ORDER BY rating DESC) rank FROM players )
	SELECT season, name, icon
			FROM ranked_players
			WHERE rank = 1 AND season < $1
			ORDER BY season`, season)
	winners := make([]models.Champion, 0)
	for rows.Next() {
		var winner models.Champion
		err = rows.Scan(&winner.Season, &winner.Name, &winner.Icon)
		if err != nil {
			return nil, err
		}
		winners = append(winners, winner)
	}
	return winners, err
}
