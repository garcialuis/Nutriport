package nutriportclient

import (
	"github.com/garcialuis/Nutriport/sdk/client/bmi"
	"github.com/garcialuis/Nutriport/sdk/client/fooditem"
	"github.com/garcialuis/Nutriport/sdk/client/tee"
	"github.com/garcialuis/Nutriport/sdk/models"
)

type NutriPortApis struct {
	FoodClient
	BMIClient
	TEEClient
}

type BMIClient interface {
	CalculateImperialBMI(weight, height float64) models.Person
}

type TEEClient interface {
	CalculateTotalEnergyExpenditure(age int, gender int, weight float64, activityLevel string) models.Person
}

type FoodClient interface {
	CreateFoodItem(foodItem models.FoodItem) models.FoodItem
	GetAllFoodItems() []models.FoodItem
	DeleteFoodItem(foodItemName string) int
	GetFoodItemByName(foodItemName string) models.FoodItem
}

func NewClient() *NutriPortApis {

	cli := new(NutriPortApis)
	cli.FoodClient = fooditem.NewClientService()
	cli.BMIClient = bmi.NewBMIService()
	cli.TEEClient = tee.NewClientService()

	return cli
}
