package utils

import (
	"net/http"
	"strconv"

	"github.com/thalkz/kart/config"
)

func ParseSeason(cfg *config.Config, r *http.Request) int {
	seasonStr := r.FormValue("season")
	season, err := strconv.Atoi(seasonStr)
	if seasonStr == "" || err != nil {
		season = cfg.GetSeason()
	}
	return season
}
