package models

type FoodItem struct {
	ID          uint32    `json:"id" gorm:"primary_key;auto_increment"`
	Name        string    `json:"name" gorm:"size:255;not null"`
	Variety     string    `json:"variety" gorm:"size:255; not null"`
	CupQuantity uint8     `json:"cupQuantity" gorm:"not null"`
	GramWeight  uint8     `json:"GMWt" gorm:"not null"`
	OnceWeight  uint8     `json:"OzWt" gorm:"not null"`
	Group       FoodGroup `json:"foodGroup" gorm:"foreignkey:FoodGroupID"`
	FoodGroupID uint8
	CarbLevel   CarbLevel `json:"carbLevel" gorm:"foreignkey:CarbLevelID"`
	CarbLevelID uint8
}
