package fooditem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/garcialuis/Nutriport/sdk/models"
)

var hostUrl string

type ClientService struct {
}

func NewClientService() *ClientService {
	initFoodItemServiceUrl()
	return &ClientService{}
}

func (service *ClientService) CreateFoodItem(foodItem models.FoodItem) models.FoodItem {

	jsonBody, err := json.Marshal(foodItem)
	if err != nil {
		log.Println("Error creating foodItem, cannot marshal input: ", err.Error())
		return models.FoodItem{}
	}

	requestBody := bytes.NewReader(jsonBody)

	url := fmt.Sprint(hostUrl, "fooditem")
	resp, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		log.Println("Unable to complete request due to: ", err.Error())
		return models.FoodItem{}
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)

	newFoodItem := models.FoodItem{}
	err = json.Unmarshal(respBytes, &newFoodItem)
	if err != nil {

	}

	return newFoodItem
}

// GetAllFoodItems retrieves all foodItems that are stored in the Nutriport Database
func (service *ClientService) GetAllFoodItems() []models.FoodItem {

	foodItems := []models.FoodItem{}

	url := fmt.Sprint(hostUrl, "fooditem")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &foodItems)
	if err != nil {
		log.Fatalln(err)
	}

	return foodItems
}

func (service *ClientService) GetFoodItemByName(foodItemName string) models.FoodItem {

	foodItem := models.FoodItem{}

	reqUrl := fmt.Sprint(hostUrl, "fooditem/")
	base, err := url.Parse(reqUrl)
	if err != nil {
		log.Println("Unable to complete request due to: ", err.Error())
		return foodItem
	}

	base.Path += foodItemName

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println("Unable to complete request due to: ", err.Error())
		return foodItem
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &foodItem)
	if err != nil {
		log.Println("Unable to unmarshall data retrieved, ", err.Error())
		return foodItem
	}

	return foodItem
}

func (service *ClientService) DeleteFoodItem(foodItemName string) int {

	client := &http.Client{}

	// foodItem := models.FoodItem{}
	reqUrl := fmt.Sprint(hostUrl, "fooditem/")
	base, err := url.Parse(reqUrl)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	base.Path += foodItemName

	req, err := http.NewRequest("DELETE", base.String(), nil)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Display Results
	fmt.Println("Response Status :", resp.Status)

	statusCode := resp.StatusCode

	if statusCode != http.StatusNoContent {
		fmt.Println("Response Body : ", string(respBody))
		return 0
	}

	return 1
}

func initFoodItemServiceUrl() {
	hostUrl = os.Getenv("SERVICE_URL")
	if len(hostUrl) == 0 {
		hostUrl = "http://localhost:8085/"
	}
}
