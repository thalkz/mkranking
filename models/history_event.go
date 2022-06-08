package models

type HistoryEvent struct {
	RaceId    int     `json:"race_id"`
	Date      string  `json:"date"`
	NewRating float64 `json:"new_rating"`
}
