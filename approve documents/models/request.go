package models

type Request struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	User   int    `json:"user"`
	Assign string `json:"assign"`
}
