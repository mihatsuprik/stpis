package models

type Groups struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
