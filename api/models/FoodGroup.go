package models

type FoodGroup struct {
	ID uint8 `json:"id" gorm:"primary_key:auto_increment`

	FoodGroupName string `json:"foodGroupName" gorm:"size:255:not null;unique"`
}
