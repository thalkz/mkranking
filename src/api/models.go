package api

type Player struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}
