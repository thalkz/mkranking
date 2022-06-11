package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/thalkz/kart/internal/config"
	"github.com/thalkz/kart/internal/database"
)

type chooseNameAndIconPage struct {
	NamePlaceholder string
	Icons           []int
}

func NewPlayerHandler(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	icon := r.FormValue("icon")
	if name != "" && icon != "" {
		return createPlayerHandler(w, r, name, icon)
	} else {
		return chooseNameAndIconHandler(w, r)
	}
}

func createPlayerHandler(w http.ResponseWriter, r *http.Request, name string, iconStr string) error {
	icon, err := strconv.Atoi(iconStr)
	if err != nil {
		return fmt.Errorf("failed to parse icon: %w", err)
	}

	id, err := database.CreatePlayer(name, config.InitialRating, icon)
	if err != nil {
		return fmt.Errorf("failed to create user %v: %w", name, err)
	}
	log.Printf("Created player %v with id=%v\n", name, id)

	http.Redirect(w, r, fmt.Sprintf("/welcome?name=%v&icon=%v", name, icon), http.StatusFound)
	return nil
}

func chooseNameAndIconHandler(w http.ResponseWriter, r *http.Request) error {
	icons := make([]int, 49)
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
