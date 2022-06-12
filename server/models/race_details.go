package models

type RaceDetails struct {
	Race
	Results []RaceDetailsResult `json:"results"`
}

type RaceDetailsResult struct {
	RaceRank   int
	UserId     int
	Name       string
	Icon       int
	NewRating  float64
	RatingDiff float64
}
