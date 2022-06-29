package utils

import (
	"net/http"
	"strconv"

	"github.com/thalkz/kart/config"
)

func ParseSeason(r *http.Request) int {
	seasonStr := r.FormValue("season")
	season, err := strconv.Atoi(seasonStr)
	if seasonStr == "" || err != nil {
		return config.Season
	}
	return season
}
