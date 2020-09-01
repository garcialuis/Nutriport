package models

type CarbLevel struct {
	// The id represents the carb level id: 1 = low, 2 = moderate, 3 = high in carbs
	ID uint8 `json:"id" gorm:"primary_key:auto_increment"`
	// The description is a string describing the level, this will be either low, moderate, or high
	Description string `json:"description" gorm:"size:255;not null"`
}
