package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/config"
	"github.com/thalkz/kart/database"
)

type chooseNameAndIconPage struct {
	NamePlaceholder string
	Icons           []int
}

func NewPlayerHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	icon := r.FormValue("icon")
	if name != "" && icon != "" {
		return createPlayerHandler(cfg, w, r, name, icon)
	} else {
		return chooseNameAndIconHandler(cfg, w, r)
	}
}

func createPlayerHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request, name string, iconStr string) error {
	icon, err := strconv.Atoi(iconStr)
	if err != nil {
		return fmt.Errorf("failed to parse icon: %w", err)
	}

	id, err := database.CreatePlayer(name, cfg.InitialRating, icon, cfg.GetSeason())
	if err != nil {
		return fmt.Errorf("failed to create user %v: %w", name, err)
	}
	log.Printf("Created player %v with id=%v\n", name, id)

	http.Redirect(w, r, fmt.Sprintf("/welcome?name=%v&icon=%v", name, icon), http.StatusFound)
	return nil
}

func chooseNameAndIconHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request) error {
	icons := make([]int, 52)
	for i := range icons {
		icons[i] = i + 1
	}

	data := chooseNameAndIconPage{
		NamePlaceholder: "Jean-Yves",
		Icons:           icons,
	}
	renderTemplate(w, "player_new.html", data)
	return nil
}
