package seed

import (
	"log"

	"github.com/garcialuis/Nutriport/api/models"
	"github.com/jinzhu/gorm"
)

var carbohydrates = []models.CarbLevel{
	models.CarbLevel{
		Description: "lowest carbohydrate",
	},
	models.CarbLevel{
		Description: "moderate carbohydrate",
	},
	models.CarbLevel{
		Description: "highest carbohydrate",
	},
}

var foodVarieties = []models.FoodVariety{
	models.FoodVariety{
		FoodVarietyName: "Raw",
	},
	models.FoodVariety{
		FoodVarietyName: "Coocked from frozen",
	},
	models.FoodVariety{
		FoodVarietyName: "Coocked from frozen with skin",
	},
	models.FoodVariety{
		FoodVarietyName: "Coocked from fresh",
	},
	models.FoodVariety{
		FoodVarietyName: "Coocked from fresh with skin",
	},
	models.FoodVariety{
		FoodVarietyName: "Coocked from canned",
	},
	models.FoodVariety{
		FoodVarietyName: "Canned",
	},
	models.FoodVariety{
		FoodVarietyName: "Not Applicable",
	},
}

var foodGroups = []models.FoodGroup{
	models.FoodGroup{
		FoodGroupName: "Vegetables",
	},
	models.FoodGroup{
		FoodGroupName: "Fuits",
	},
	models.FoodGroup{
		FoodGroupName: "Grains",
	},
	models.FoodGroup{
		FoodGroupName: "Protein",
	},
	models.FoodGroup{
		FoodGroupName: "Dairy",
	},
}

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
}
