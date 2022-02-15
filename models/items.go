package models

// Item items Struct (Model)
type Item struct {
	Id        string `gorm:"type:varchar(255)"`
	Text      string `json:"text" gorm:"type:varchar(255)"`
	Completed bool   `json:"completed" gorm:"type:bool"`
	Datetime  string `json:"datetime" gorm:"type:varchar(255)"`
}
