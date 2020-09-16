package controllers

import (
	"errors"
	"net/http"

	"github.com/garcialuis/Nutriport/api/responses"
)

// TEE = RMR x activity factor

/*
Men:

Activity levels:
Sedentary: 0.3
Lightly Active: 0.6
Moderately Active: 0.7
Very Active: 1.1
Extremely Active: 1.4

RMR Equation based on age:

0-3 (60.9 x weight) - 54
3-10 (22.7 x weight) + 495
10-18 (17.5 x weight) + 651
18-30 (15.3 x weight) + 679
30-60 (11.6 x weight) + 879
>60 (13.5 x weight) + 487
*/

/*
Women:

Activity levels:
Sedentary: 0.3
Lightly Active: 0.5
Moderately Active: 0.6
Very Active: 0.9
Extremely Active: 1.2

RMR Equation based on age:

0-3 (61.0 x weight) - 51
3-10 (22.5 x weight) + 499
10-18 (12.2 x weight) + 746
18-30 (14.7 x weight) + 496
30-60 (8.7 x weight) + 829
>60 (10.5 x weight) + 596
*/

var rmrValues = map[int][2]float64{
	0:  {60.9, -54},
	1:  {22.7, 495},
	2:  {17.5, 651},
	3:  {15.3, 679},
	4:  {11.6, 879},
	5:  {13.5, 487},
	6:  {61.0, -51},
	7:  {22.5, 499},
	8:  {12.2, 746},
	9:  {14.7, 496},
	10: {8.7, 829},
	11: {10.5, 596},
}

var activityLevels = map[string][2]float64{
	"sedentary":         {0.3, 0.3},
	"lightly active":    {0.6, 0.5},
	"moderately active": {0.7, 0.6},
	"very active":       {1.1, 0.9},
	"extremely active":  {1.4, 1.2},
}

func (server *Server) GetTotalEnergyExpenditure(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	activityLevelParams, ok := vars["activitylevel"]

	if !ok || len(activityLevelParams[0]) < 1 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Url Param 'height' is missing"))
		return
	}

	activityLevel := activityLevelParams[0]

	// personInfo := models.Person{
	// 	Height: height,
	// 	Weight: weight,
	// }

	tee := calculateTEE(25, 0, 143, activityLevel)

	responses.JSON(w, http.StatusOK, tee)

	// responses.JSON(w, http.StatusOK, personInfo)
}

func calculateTEE(age int, gender int, weight float64, activityLevel string) (tee float64) {

	rmr := getRestingMetabollicRate(age, gender, weight)
	activityLevelFactor := activityLevels[activityLevel][gender]

	tee = rmr * activityLevelFactor

	return tee
}

func getRestingMetabollicRate(age int, gender int, weight float64) (rmr float64) {

	rmrIndex := getRMRPairIndex(age, gender)
	factor, addend := getRMRValues(rmrIndex)

	rmr = (factor * weight) + addend

	return rmr
}

func getActivityLevel(activityLevel string, gender int) float64 {
	return activityLevels[activityLevel][gender]
}

func getRMRPairIndex(age int, gender int) (pairIndex int) {

	if age < 3 {
		pairIndex = 0
	} else if age >= 3 && age < 10 {
		pairIndex = 1
	} else if age >= 10 && age < 18 {
		pairIndex = 2
	} else if age >= 18 && age < 30 {
		pairIndex = 3
	} else if age >= 30 && age < 60 {
		pairIndex = 4
	} else if age >= 60 {
		pairIndex = 5
	}

	pairIndex = (6 * gender) + pairIndex

	return pairIndex
}

func getRMRValues(index int) (factor float64, addends float64) {

	rmrEquationValues := rmrValues[index]

	return rmrEquationValues[0], rmrEquationValues[1]
}
