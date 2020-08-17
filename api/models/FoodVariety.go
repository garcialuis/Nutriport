package models

type FoodVariety struct {
	ID              uint16 `json:"ID" gorm:"primary_key;auto_increment"`
	FoodVarietyName string `json:"variety" gorm:"size:255; not null"`
}
