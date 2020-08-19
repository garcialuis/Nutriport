package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type FoodItem struct {
	ID            uint32      `json:"id" gorm:"primary_key;auto_increment"`
	Name          string      `json:"name" gorm:"size:255;not null"`
	Variety       FoodVariety `json:"variety" gorm:"association_autoupdate:false;association_autocreate:false;foreignkey:FoodVarietyID"`
	FoodVarietyID uint16
	CupQuantity   float32   `json:"cupQuantity" gorm:"not null"`
	GramWeight    float32   `json:"GMWt" gorm:"not null"`
	OnceWeight    float32   `json:"OzWt" gorm:"not null"`
	Group         FoodGroup `json:"foodGroup" gorm:"association_autoupdate:false;association_autocreate:false;foreignkey:FoodGroupID"`
	FoodGroupID   uint8
	CarbLevel     CarbLevel `json:"carbLevel" gorm:"association_autoupdate:false;association_autocreate:false;foreignkey:CarbLevelID"`
	CarbLevelID   uint8
}

func (item *FoodItem) SelectByName(db *gorm.DB, name string) (*FoodItem, error) {

	nameStr := "%" + name + "%"
	fmt.Printf("nameStr for query: %s", nameStr)

	err := db.Debug().Model(&FoodItem{}).Where("name LIKE ?", nameStr).Take(&item).Error
	if err != nil {
		return &FoodItem{}, err
	}

	return item, nil
}

func (item *FoodItem) SelectByID(db *gorm.DB, foodItemID uint32) (*FoodItem, error) {

	err := db.Debug().Model(&FoodItem{}).Where("name = ?", foodItemID).Take(&item).Error
	if err != nil {
		return &FoodItem{}, err
	}

	return item, nil
}

func (item *FoodItem) SaveFoodItem(db *gorm.DB) (*FoodItem, error) {

	err := db.Debug().Model(&FoodItem{}).Create(&item).Error //.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).
	if err != nil {
		return &FoodItem{}, err
	}

	return item, nil
}
