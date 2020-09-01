package models

type FoodGroup struct {
	// ID is the database id that is given to the food group record
	// This value is required when inputing a new food item.
	ID uint8 `json:"id"`
	// FoodGroupName is a string describing the food group
	FoodGroupName string `json:"foodGroupName"`
}
