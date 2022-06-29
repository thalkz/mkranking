package models

type Race struct {
	Id      int           `json:"id"`
	Date    string        `json:"date"`
	Results [4]RaceResult `json:"results"`
}

type RaceResult struct {
	Rank       int     `json:"rank"`
	UserId     int     `json:"user_id"`
	Name       string  `json:"name"`
	Icon       int     `json:"icon"`
	NewRating  float64 `json:"new_rating"`
	RatingDiff float64 `json:"rating_diff"`
}
