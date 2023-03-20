package models

type Approve struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Requestid string `json:"requestid"`
	Docid     int    `json:"docid"`
}
