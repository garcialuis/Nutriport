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

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, foodItemCreated.ID))
	responses.JSON(w, http.StatusCreated, foodItemCreated)
}

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

func (server *Server) GetAllFoodItems(w http.ResponseWriter, r *http.Request) {

	foodItem := models.FoodItem{}

	foodItems, err := foodItem.SelectAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, foodItems)
}
