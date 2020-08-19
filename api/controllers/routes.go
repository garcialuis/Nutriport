package controllers

import "github.com/garcialuis/Nutriport/api/middlewares"

func (s *Server) InitializeRoutes() {
	// Home
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// FoodItems
	s.Router.HandleFunc("/fooditem/{foodName}", middlewares.SetMiddlewareJSON(s.GetFoodItemByName)).Methods("GET")
	s.Router.HandleFunc("/fooditem", middlewares.SetMiddlewareJSON(s.CreateFoodItem)).Methods("POST")
}
