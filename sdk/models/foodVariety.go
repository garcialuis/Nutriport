package models

type FoodVariety struct {
	// FoodVarietyName is a description of the variety
	// For example, a variety can be coocked from fresh or frozen
	FoodVarietyName string `json:"variety,omitempty"`
	// ID is the database id assigned to the food group
	// This id is used to referece the variety from the food items table
	ID *uint16 `json:"ID"`
}
