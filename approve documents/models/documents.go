package models

type Documents struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Creater int    `json:"creater"`
}
