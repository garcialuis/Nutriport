// Package classification Nutriport API
//
// Documentation for Nutriport API
//
//		Schemes: http
//		BasePath: /
//		Title: NutriportAPI
//		Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
// swagger:meta
package controllers

import "github.com/garcialuis/Nutriport/api/models"

// A single food item record
// swagger:response foodItemResponse
type foodItemResponseWrapper struct {
	// Single food item
	// in: body
	Body models.FoodItem
}

// A list of food items
// swagger:response foodItemsResponse
type foodItemsResponseWrapper struct {
	// All foodItems currently in db
	// in: body
	Body []models.FoodItem
}

// foodName is the name of the food
// swagger:parameters GetFoodItemByName DeleteFoodItemByName
type foodNameParamWrapper struct {
	// The name of a food item
	// in: path
	// required: true
	FoodName string `json:"foodName"`
}
