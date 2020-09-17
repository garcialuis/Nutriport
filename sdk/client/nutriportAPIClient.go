package client

import (
	"fmt"

	"github.com/garcialuis/Nutriport/sdk/client/bmi"
	"github.com/garcialuis/Nutriport/sdk/client/fooditem"
	"github.com/garcialuis/Nutriport/sdk/client/tee"
)

type NutriPortApis struct {
	FoodClient *fooditem.ClientService
	BMIClient  *bmi.BMIClientService
	TEEClient  *tee.ClientService
}

func New() *NutriPortApis {

	cli := new(NutriPortApis)
	cli.FoodClient = fooditem.NewClientService()
	cli.BMIClient = bmi.NewBMIService()
	cli.TEEClient = tee.NewClientService()

	return cli
}

func NewTestFunc() {
	npa := New()
	info := npa.BMIClient.CalculateImperialBMI(12.0, 1234.5)
	fmt.Println(info)
}
