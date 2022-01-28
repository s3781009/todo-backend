package models

// Items init item var as a slice

// Item items Struct (Model)
type Item struct {
	Id        string
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
	Datetime  string `json:"datetime"`
}
