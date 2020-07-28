package datamodels

type Song struct {
	ID		int64  `json:"id"`
	Name	string `json:"name"`
	Link	string    `json:"link"`
	Time	string `json:"time"`
	Poster 	string `json:"poster"`
}