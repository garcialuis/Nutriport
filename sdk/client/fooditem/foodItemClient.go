package fooditem

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/garcialuis/Nutriport/sdk/models"
)

type ClientService struct {
}

func NewClientService() *ClientService {
	return &ClientService{}
}

func (service *ClientService) CreateFoodItem(foodItem *models.FoodItem) models.FoodItem {

	jsonBody, err := json.Marshal(foodItem)
	if err != nil {
		// Unable to marshall input
	}

	requestBody := bytes.NewReader(jsonBody)

	resp, err := http.Post("", "application/json", requestBody)

	if err != nil {
		//
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

	resp, err := http.Get("http://localhost:8085/fooditem")
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

	base, err := url.Parse("")
	if err != nil {

	}

	base.Path += foodItemName

	resp, err := http.Get(base.String())
	if err != nil {

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &foodItem)
	if err != nil {

	}

	return foodItem
}
