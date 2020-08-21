package models

// FoodVariety represents the variety for a food item
// For example, a variety for a brocolli can be coocked from fresh or frozen
//
// swagger:model
type FoodVariety struct {
	// ID is the database id assigned to the food group
	// This id is used to referece the variety from the food items table
	// required: true
	// min: 1
	ID uint16 `json:"ID" gorm:"primary_key;auto_increment"`
	// FoodVarietyName is a description of the variety
	// For example, a variety can be coocked from fresh or frozen
	// required: false
	FoodVarietyName string `json:"variety" gorm:"size:255; not null"`
}
