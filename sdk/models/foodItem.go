package models

type FoodItem struct {
	// ID is the database assigned id to the food item record
	ID uint32 `json:"id"`
	// Name is the name of the food item
	Name string `json:"name"`
	// Variety is the state of the foodItem at the time before being eaten
	Variety FoodVariety `json:"variety"`
	// FoodVarietyID is a reference to the FoodVariety record in the db
	FoodVarietyID uint16
	// CupQuantity represents the amount that is allowed to intake in Cup(s)
	CupQuantity float32 `json:"cupQuantity"`
	// GramWeight represents the allowed weight to intake - in grams
	GramWeight float32 `json:"GMWt"`
	// OnceWeight represents the allowed weight to intake - in ounces
	OnceWeight float32 `json:"OzWt"`
	// Group holds information about the food groups that the food item belongs to
	Group FoodGroup `json:"foodGroup"`
	// FoodGroupID is the reference to the group from the food group db table
	FoodGroupID uint8
	// CarbLevel is the information about what level of carbohydrates can be found in the item
	CarbLevel CarbLevel `json:"carbLevel"`
	// CarbLevelID is the reference to the carb level in the database table
	CarbLevelID uint8
}
