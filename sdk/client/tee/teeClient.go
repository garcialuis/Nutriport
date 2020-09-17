package tee

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/garcialuis/Nutriport/sdk/models"
)

type ClientService struct {
}

func NewClientService() *ClientService {
	return &ClientService{}
}

func (service *ClientService) CalculateTotalEnergyExpenditure(age int, gender int, weight float64, activityLevel string) models.Person {
	personInfo := models.Person{}
	client := &http.Client{}

	ageStr := fmt.Sprintf("%f", age)
	genderStr := fmt.Sprintf("%f", gender)
	weightStr := fmt.Sprintf("%f", weight)

	req, err := http.NewRequest("GET", "http://localhost:8085/tee", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("age", ageStr)
	q.Add("gender", genderStr)
	q.Add("weight", weightStr)
	q.Add("activitylevel", activityLevel)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error sending request to server: ", err.Error())
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
