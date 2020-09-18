package tests

import (
	"fmt"
	"testing"
	"time"

	nutriportclient "github.com/garcialuis/Nutriport/sdk/client"
	"github.com/stretchr/testify/assert"
)

func TestCalculateImperialBMI(t *testing.T) {

	time.Sleep(2 * time.Second)

	nutriportClient := nutriportclient.NewClient()

	weight := 160.0
	height := 67.0

	personInfo := nutriportClient.CalculateImperialBMI(weight, height)

	fmt.Println("BMI: ", personInfo.BMI, personInfo.BMIDescription)

	fmt.Println(personInfo)

	assert.Equal(t, 160.0, personInfo.Weight)
	assert.Equal(t, 67.0, personInfo.Height)
	assert.Equal(t, 25.1, personInfo.BMI)
	assert.Equal(t, "Overweight", personInfo.BMIDescription)
}
