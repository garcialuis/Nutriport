package tests

import (
	"testing"

	nutriportclient "github.com/garcialuis/Nutriport/sdk/client"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalEnergyExpenditureMale(t *testing.T) {

	nutriportClient := nutriportclient.NewClient()
	personInfo := nutriportClient.TEEClient.CalculateTotalEnergyExpenditure(25, 0, 143, "moderately active")

	assert.Equal(t, 2006.83, personInfo.TEE)
}

func TestCalculateTotalEnergyExpenditureFemale(t *testing.T) {

	nutriportClient := nutriportclient.NewClient()
	personInfo := nutriportClient.TEEClient.CalculateTotalEnergyExpenditure(25, 1, 130, "extremely active")

	assert.Equal(t, 2888.40, personInfo.TEE)
}

func TestCalculateTotalEnergyExpenditureFemale2(t *testing.T) {

	nutriportClient := nutriportclient.NewClient()
	personInfo := nutriportClient.CalculateTotalEnergyExpenditure(42, 1, 122.53, "sedentary")

	assert.Equal(t, 568.50, personInfo.TEE)
}
