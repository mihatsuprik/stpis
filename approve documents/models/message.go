package models

type Message struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Requestid int  `json:"requestid"`
	Userid    int  `json:"userid"`
	Notice    int  `json:"notice"`
}