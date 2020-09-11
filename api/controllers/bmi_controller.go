package controllers

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/garcialuis/Nutriport/api/models"
	"github.com/garcialuis/Nutriport/api/responses"
)

func (server *Server) GetMetricBMI(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	heightParams, ok := vars["height"]

	if !ok || len(heightParams[0]) < 1 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Url Param 'height' is missing"))
		return
	}

	height, err := strconv.ParseFloat(heightParams[0], 64)

	if err != nil || height < 1 {
		if err == nil {
			err = errors.New("Invalid 'height' specified, must be > 0")
		}
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	weightParams, ok := vars["weight"]

	if !ok || len(weightParams[0]) < 1 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Url Param 'weight' is missing"))
		return
	}

	weight, err := strconv.ParseFloat(weightParams[0], 64)
	if err != nil || weight < 1 {
		if err == nil {
			err = errors.New("Invalid 'weight' specified, must be > 0")
		}
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	personInfo := models.Person{
		Height: height,
		Weight: weight,
	}

	bmi, bmiDescription := calculateMetricBMI(weight, height)
	personInfo.BMI = bmi
	personInfo.BMIDescription = bmiDescription

	responses.JSON(w, http.StatusOK, personInfo)

}

func (server *Server) GetImperialBMI(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	heightParams, ok := vars["height"]

	if !ok || len(heightParams[0]) < 1 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Url Param 'height' is missing"))
		return
	}

	height, err := strconv.ParseFloat(heightParams[0], 64)

	if err != nil || height < 1 {
		if err == nil {
			err = errors.New("Invalid 'height' specified, must be > 0")
		}
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	weightParams, ok := vars["weight"]

	if !ok || len(weightParams[0]) < 1 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Url Param 'weight' is missing"))
		return
	}

	weight, err := strconv.ParseFloat(weightParams[0], 64)
	if err != nil || weight < 1 {
		if err == nil {
			err = errors.New("Invalid 'weight' specified, must be > 0")
		}
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	personInfo := models.Person{
		Height: height,
		Weight: weight,
	}

	bmi, bmiDescription := calculateImperialBMI(weight, height)
	personInfo.BMI = bmi
	personInfo.BMIDescription = bmiDescription

	responses.JSON(w, http.StatusOK, personInfo)
}

func calculateMetricBMI(weight, height float64) (bmi float64, bmiDescription string) {

	bmi = weight / math.Pow(height, 2)
	bmi = math.Round(bmi*10) / 10
	bmiDescription = getBMIDescription(bmi)

	return bmi, bmiDescription
}

func calculateImperialBMI(weight, height float64) (bmi float64, bmiDescription string) {

	bmi = 703 * weight / math.Pow(height, 2)
	bmi = math.Round(bmi*10) / 10
	bmiDescription = getBMIDescription(bmi)

	return bmi, bmiDescription
}

func getBMIDescription(bmi float64) (bmiDescription string) {

	if bmi >= 30 {
		bmiDescription = "Obese"
	} else if 25 <= bmi && bmi <= 29.9 {
		bmiDescription = "Overweight"
	} else if 18.6 <= bmi && bmi <= 24.9 {
		bmiDescription = "Healthy"
	} else if bmi <= 18.5 {
		bmiDescription = "Thin"
	}

	return bmiDescription
}
