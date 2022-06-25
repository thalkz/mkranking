package models

type Race struct {
	Id      int           `json:"id"`
	Date    string        `json:"date"`
	Results [4]RaceResult `json:"results"`
}

type RaceResult struct {
	Rank       int
	UserId     int
	Name       string
	Icon       int
	NewRating  float64
	RatingDiff float64
}
