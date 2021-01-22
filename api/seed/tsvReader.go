package seed

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/garcialuis/Nutriport/sdk/models"
)

var foodItemList []models.FoodItem

func getFoodItemList() []models.FoodItem {
	// Read List from tsv file
	csvReadAll()

	return foodItemList
}

func csvReadAll() {
	recordFile, err := os.Open("../../resources/vegetable_conversion.tsv")
	if err != nil {
		fmt.Println("An error encountered :", err)
	}

	reader := csv.NewReader(recordFile)
	reader.Comma = '\t'
	reader.Comment = '#'

	records, _ := reader.ReadAll()

	for _, record := range records {
		foodItem := constructFoodItem(record)
		addFoodItem(foodItem)
	}
}

func addFoodItem(foodItem models.FoodItem) {
	foodItemList = append(foodItemList, foodItem)
}

func constructFoodItem(record []string) models.FoodItem {

	name := record[0]
	variety, varietyID := getFoodVariety(record[1])
	cupQuantity, _ := strconv.ParseFloat(record[2], 32)
	gramQuantity, _ := strconv.ParseFloat(record[3], 32)
	ounceQuantity, _ := strconv.ParseFloat(record[4], 32)
	carbLevelID, _ := strconv.Atoi(record[5])
	carbDescription := getCarbLevelDescription(carbLevelID)

	foodItem := models.FoodItem{
		Name:          name,
		Variety:       models.FoodVariety{ID: &varietyID, FoodVarietyName: variety},
		FoodVarietyID: varietyID,
		CupQuantity:   float32(cupQuantity),
		GramWeight:    float32(gramQuantity),
		OnceWeight:    float32(ounceQuantity),
		Group:         models.FoodGroup{ID: 1, FoodGroupName: "Vegetables"},
		FoodGroupID:   1,
		CarbLevel:     models.CarbLevel{ID: uint8(carbLevelID), Description: carbDescription},
		CarbLevelID:   uint8(carbLevelID),
	}

	return foodItem
}

func getCarbLevelDescription(carbLevelID int) string {

	carbLevels := [3]string{"lowest carbohydrate", "moderate carbohydrate", "highest carbohydrate"}

	return carbLevels[carbLevelID-1]
}

func getFoodVariety(variety string) (string, uint16) {

	foodVarietiesMap := map[string]uint16{
		"RAW":                          1,
		"COOKED FROM FROZEN":           2,
		"COOKED FROM FROZEN WITH SKIN": 3,
		"COOKED FROM FRESH":            4,
		"COOKED FROM FRESH WITH SKIN":  5,
		"COOKED FROM CANNED":           6,
		"CANNED":                       7,
		"N/A":                          8,
	}

	varietyID := foodVarietiesMap[strings.ToUpper(variety)]

	return variety, varietyID
}
