package seed

import (
	"log"

	"github.com/garcialuis/Nutriport/api/models"
	"github.com/jinzhu/gorm"
)

var carbohydrates = []models.CarbLevel{
	{
		Description: "lowest carbohydrate",
	},
	{
		Description: "moderate carbohydrate",
	},
	{
		Description: "highest carbohydrate",
	},
}

var foodVarieties = []models.FoodVariety{
	{
		FoodVarietyName: "Raw",
	},
	{
		FoodVarietyName: "Cooked from frozen",
	},
	{
		FoodVarietyName: "Cooked from frozen with skin",
	},
	{
		FoodVarietyName: "Cooked from fresh",
	},
	{
		FoodVarietyName: "Cooked from fresh with skin",
	},
	{
		FoodVarietyName: "Cooked from canned",
	},
	{
		FoodVarietyName: "Canned",
	},
	{
		FoodVarietyName: "Not Applicable",
	},
}

var foodGroups = []models.FoodGroup{
	{
		FoodGroupName: "Vegetables",
	},
	{
		FoodGroupName: "Fuits",
	},
	{
		FoodGroupName: "Grains",
	},
	{
		FoodGroupName: "Protein",
	},
	{
		FoodGroupName: "Dairy",
	},
}

// Load takes care of dropping the database tables if the exist and creating new tables to seed the initial data in.
// This function should only be used in development environments.
// ** THIS IS NOT FOR PRODUCTION PURPOSES **
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.FoodItem{}, &models.CarbLevel{}, &models.FoodGroup{}, &models.FoodVariety{}).Error

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.FoodVariety{}, &models.CarbLevel{}, &models.FoodGroup{}, &models.FoodItem{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table(s): %v", err)
	}

	err = db.Debug().Model(&models.FoodItem{}).AddForeignKey("food_variety_id", "food_varieties(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.FoodItem{}).AddForeignKey("carb_level_id", "carb_levels(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.FoodItem{}).AddForeignKey("food_group_id", "food_groups(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range carbohydrates {
		err = db.Debug().Model(&models.CarbLevel{}).Create(&carbohydrates[i]).Error
		if err != nil {
			log.Fatalf("Could not seed table with Carbohydrate levels: %v", err)
		}
	}

	for i := range foodVarieties {
		err = db.Debug().Model(&models.FoodVariety{}).Create(&foodVarieties[i]).Error
		if err != nil {
			log.Fatalf("Could not seed FoodVarieties table: %v", err)
		}
	}

	for i := range foodGroups {
		err = db.Debug().Model(&models.FoodGroup{}).Create(&foodGroups[i]).Error
		if err != nil {
			log.Fatalf("Could not seed FoodGroup table: %v", err)
		}
	}

	foodItemList := getFoodItemList()

	for i := range foodItemList {
		err = db.Model(&models.FoodItem{}).Create(&foodItemList[i]).Error
		if err != nil {
			log.Fatalf("Could not seed foodItems from tsv file at index: %d : %v", i, err)
		}
	}
}
