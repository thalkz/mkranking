package models

type Player struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Rank   int     `json:"rank"`
	Icon   int     `json:"icon"`
	Rating float64 `json:"rating"`
}
