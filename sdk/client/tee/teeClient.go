package tee

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/garcialuis/Nutriport/sdk/models"
)

var hostUrl string

type ClientService struct {
}

func NewClientService() *ClientService {
	initTEEServiceUrl()
	return &ClientService{}
}

func (service *ClientService) CalculateTotalEnergyExpenditure(age int, gender int, weight float64, activityLevel string) models.Person {
	personInfo := models.Person{}
	client := &http.Client{}

	ageStr := strconv.Itoa(age)
	genderStr := strconv.Itoa(gender)
	weightStr := fmt.Sprintf("%f", weight)

	url := fmt.Sprint(hostUrl, "tee")
	req, err := http.NewRequest("GET", url, nil)
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

func initTEEServiceUrl() {
	hostUrl = os.Getenv("NUTRIPORT_SERVICE_URL")
	if len(hostUrl) == 0 {
		hostUrl = "http://localhost:8085/"
	}
}
