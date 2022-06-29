package models

type History struct {
	Players []Player               `json:"players"`
	Events  map[int][]HistoryEvent `json:"events"` // Map of userId -> []HistortEvent
}

type HistoryEvent struct {
	RaceId int     `json:"race_id"`
	Rating float64 `json:"rating"`
}
