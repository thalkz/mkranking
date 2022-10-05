package web

import (
	"fmt"
	"net/http"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
	"github.com/thalkz/kart/models"
)

type ChampionsPage struct {
	Champions []models.Champion
}

func ChampionsHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	champions, err := database.GetChampions(cfg.GetSeason())
	if err != nil {
		return fmt.Errorf("failed to get winners: %w", err)
	}

	for i := range champions {
		champions[i].Date = cfg.GetStartDate(champions[i].Season).Format("Jan 2006")
	}

	data := ChampionsPage{
		Champions: champions,
	}
	renderTemplate(w, "champions.html", data)
	return nil
}
