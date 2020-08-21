package models

// FoodGroup identifies the food group that a food item belongs to
// Food groups can be: fruit, dairy, protein, etc...
//
// swagger:model
type FoodGroup struct {
	// ID is the database id that is given to the food group record
	// This value is required when inputing a new food item.
	// required: true
	// min: 1
	ID uint8 `json:"id" gorm:"primary_key:auto_increment"`
	// FoodGroupName is a string describing the food group
	// example: fruit or dairy
	// required: false
	FoodGroupName string `json:"foodGroupName" gorm:"size:255:not null;unique"`
}
