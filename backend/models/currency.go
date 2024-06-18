package models

type Currency struct {
	ID           uint   `gorm:"primaryKey" json:"Cur_ID"`
	Abbreviation string `gorm:"uniqueKey" json:"Cur_Abbreviation"`
	Name         string `json:"Cur_Name"`
}
