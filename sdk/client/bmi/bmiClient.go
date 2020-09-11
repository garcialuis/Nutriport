package bmi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/garcialuis/Nutriport/sdk/models"
)

type BMIClientService struct {
}

func NewBMIService() *BMIClientService {
	return &BMIClientService{}
}

func (service *BMIClientService) CalculateImperialBMI(weight, height float64) models.Person {

	personInfo := models.Person{}
	client := &http.Client{}

	weightStr := fmt.Sprintf("%f", weight)
	heightStr := fmt.Sprintf("%f", height)

	req, err := http.NewRequest("GET", "http://localhost:8085/imperial/bmi", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("weight", weightStr)
	q.Add("height", heightStr)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error sending request to server", err.Error())
		return models.Person{}
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &personInfo)
	if err != nil {
		log.Println(err.Error())
	}

	return personInfo
}
