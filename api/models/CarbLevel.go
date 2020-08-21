package models

// CarbLevel is a numeric rank representation of the carbohydrates present in the food item
// The higher the level, the more carbohydrates the item has.
//
// swagger:model
type CarbLevel struct {
	// The id represents the carb level id: 1 = low, 2 = moderate, 3 = high in carbs
	// required: true
	// min: 1
	ID uint8 `json:"id" gorm:"primary_key:auto_increment"`
	// The description is a string describing the level, this will be either low, moderate, or high
	// required: false
	Description string `json:"description" gorm:"size:255;not null"`
}
