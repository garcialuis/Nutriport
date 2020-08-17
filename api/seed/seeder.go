package seed

import (
	"log"

	"github.com/garcialuis/Nutriport/api/models"
	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.FoodItem{}, &models.CarbLevel{}, &models.FoodGroup{}).Error

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.CarbLevel{}, &models.FoodGroup{}, &models.FoodItem{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table(s): %v", err)
	}

	err = db.Debug().Model(&models.FoodItem{}).AddForeignKey("carb_level_id", "carb_levels(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.FoodItem{}).AddForeignKey("food_group_id", "food_groups(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}
