package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/garcialuis/Nutriport/api/models"
	"github.com/garcialuis/Nutriport/api/responses"
	"github.com/gorilla/mux"
)

// CreateFoodItem is responsible for creating/storing a new food item into the db
// swagger:route POST /fooditem foodItem CreateFoodItem
//
//	Responses:
//		201: foodItemResponse
//		422: description: Unprocessable Entity
//		500: description: Internal Server Error
func (server *Server) CreateFoodItem(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	foodItem := models.FoodItem{}
	err = json.Unmarshal(body, &foodItem)
	fmt.Println("FoodItem Received: ", foodItem)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	foodItemCreated, err := foodItem.SaveFoodItem(server.DB)

	fmt.Println("Food Item saved into DB")

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, foodItemCreated.ID))
	responses.JSON(w, http.StatusCreated, foodItemCreated)
}

// GetFoodItemByName is responsible for retrieving a food item with a specified name
// swagger:route GET /fooditem/{foodName} foodItem GetFoodItemByName
//
//	Responses:
//		200: foodItemResponse
//		422: description: Unprocessable Entity
//		500: description: Internal Server Error
func (server *Server) GetFoodItemByName(w http.ResponseWriter, r *http.Request) {

	foodItem := models.FoodItem{}
	var err error

	vars := mux.Vars(r)
	foodName := vars["foodName"]

	if len(foodName) < 1 {
		err = fmt.Errorf("Invalid string length")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	itemFound, err := foodItem.SelectByName(server.DB, foodName)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, itemFound)
}

// GetAllFoodItems is responsible for retrieving all food items in the database
// swagger:route GET /fooditem foodItem GetAllFoodItems
//
//	Responses:
//		200: foodItemsResponse
//		422: description: Unprocessable Entity
//		500: description: Internal Server Error
func (server *Server) GetAllFoodItems(w http.ResponseWriter, r *http.Request) {

	foodItem := models.FoodItem{}

	foodItems, err := foodItem.SelectAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, foodItems)
}

// DeleteFoodItemByName Handler:
// swagger:route DELETE /fooditem/{foodName} foodItem DeleteFoodItemByName
//
// Responses:
//		204: description: No Content
//		400: description: Bad Request
//		404: description: FoodItem Not Found
func (server *Server) DeleteFoodItemByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	foodItemName := vars["foodName"]

	foodItem := models.FoodItem{}

	_, err := foodItem.DeleteItem(server.DB, foodItemName)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Entity", foodItemName)
	responses.JSON(w, http.StatusNoContent, "")
}
