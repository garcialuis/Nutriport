package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

// FoodItem is the model that contains all information about an item
// It includes the food variety, carb level, and food group information for a given item
//
// swagger:model
type FoodItem struct {
	// ID is the database assigned id to the food item record
	// required: false
	// min: 1
	ID uint32 `json:"id" gorm:"primary_key;auto_increment"`
	// Name is the name of the food item
	// required: true
	Name string `json:"name" gorm:"size:255;not null"`
	// Variety is the state of the foodItem at the time before being eaten
	// Example: raw, coocked from fresh, or coocked from frozen
	// requried: true
	Variety FoodVariety `json:"variety" gorm:"association_autoupdate:false;association_autocreate:false;foreignkey:FoodVarietyID"`
	// FoodVarietyID is a reference to the FoodVariety record in the db
	// required: false
	// min: 1
	FoodVarietyID uint16
	// CupQuantity represents the amount that is allowed to intake in Cup(s)
	// required: true
	CupQuantity float32 `json:"cupQuantity" gorm:"not null"`
	// GramWeight represents the allowed weight to intake - in grams
	// required: true
	GramWeight float32 `json:"GMWt" gorm:"not null"`
	// OnceWeight represents the allowed weight to intake - in ounces
	// required: true
	OnceWeight float32 `json:"OzWt" gorm:"not null"`
	// Group holds information about the food groups that the food item belongs to
	// requried: true
	Group FoodGroup `json:"foodGroup" gorm:"association_autoupdate:false;association_autocreate:false;foreignkey:FoodGroupID"`
	// FoodGroupID is the reference to the group from the food group db table
	// required: false
	// min: 1
	FoodGroupID uint8
	// CarbLevel is the information about what level of carbohydrates can be found in the item
	// required: true
	CarbLevel CarbLevel `json:"carbLevel" gorm:"association_autoupdate:false;association_autocreate:false;foreignkey:CarbLevelID"`
	// CarbLevelID is the reference to the carb level in the database table
	// required: false
	CarbLevelID uint8
}

// SelectByName is the function used to retrieve a foodItem w/a specified name
func (item *FoodItem) SelectByName(db *gorm.DB, name string) (*FoodItem, error) {

	nameStr := "%" + name + "%"
	fmt.Printf("nameStr for query: %s", nameStr)

	err := db.Debug().Model(&FoodItem{}).Where("name LIKE ?", nameStr).Take(&item).Error
	if err != nil {
		return &FoodItem{}, err
	}

	return item, nil
}

// SelectByID is the function to retrieve a foodItem by its database id
func (item *FoodItem) SelectByID(db *gorm.DB, foodItemID uint32) (*FoodItem, error) {

	err := db.Debug().Model(&FoodItem{}).Where("name = ?", foodItemID).Take(&item).Error
	if err != nil {
		return &FoodItem{}, err
	}

	return item, nil
}

// SaveFoodItem is the function to create a new foodItem record in the foodItems table
func (item *FoodItem) SaveFoodItem(db *gorm.DB) (*FoodItem, error) {

	err := db.Debug().Model(&FoodItem{}).Create(&item).Error
	if err != nil {
		return &FoodItem{}, err
	}

	return item, nil
}

// SelectAll is the function that allows us to retrieve all the foodItems in the db table
func (item *FoodItem) SelectAll(db *gorm.DB) (*[]FoodItem, error) {

	var err error

	foodItems := []FoodItem{}
	err = db.Debug().Model(&FoodItem{}).Find(&foodItems).Error

	if err != nil {
		return &[]FoodItem{}, err
	}
	if len(foodItems) > 0 {
		for i := range foodItems {
			err = db.Model(&FoodVariety{}).Where("id = ?", foodItems[i].FoodVarietyID).Take(&foodItems[i].Variety).Error
			if err != nil {
				return &[]FoodItem{}, err
			}

			err = db.Model(&FoodGroup{}).Where("id = ?", foodItems[i].FoodGroupID).Take(&foodItems[i].Group).Error
			if err != nil {
				return &[]FoodItem{}, err
			}

			err = db.Model(&CarbLevel{}).Where("id = ?", foodItems[i].CarbLevelID).Take(&foodItems[i].CarbLevel).Error
			if err != nil {
				return &[]FoodItem{}, err
			}
		}
	}

	return &foodItems, nil
}

// DeleteItem is to be used to remove an item from the food_items table with a specified name
func (item *FoodItem) DeleteItem(db *gorm.DB, foodName string) (int64, error) {

	db = db.Debug().Model(&FoodItem{}).Where("name = ?", foodName).Take(&FoodItem{}).Delete(&FoodItem{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("FoodItem specified not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
