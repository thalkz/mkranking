package models

type Race struct {
	Id      int    `json:"id"`
	Results []int  `json:"results"`
	Date    string `json:"date"`
}
