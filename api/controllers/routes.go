package controllers

import (
	"net/http"

	"github.com/garcialuis/Nutriport/api/middlewares"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) InitializeRoutes() {
	// Home
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// FoodItems
	s.Router.HandleFunc("/fooditem/{foodName}", middlewares.SetMiddlewareJSON(s.GetFoodItemByName)).Methods("GET")
	s.Router.HandleFunc("/fooditem", middlewares.SetMiddlewareJSON(s.CreateFoodItem)).Methods("POST")
	s.Router.HandleFunc("/fooditem", middlewares.SetMiddlewareJSON(s.GetAllFoodItems)).Methods("GET")
	s.Router.HandleFunc("/fooditem/{foodName}", middlewares.SetMiddlewareJSON(s.DeleteFoodItemByName)).Methods("DELETE")

	// BMI
	s.Router.HandleFunc("/metric/bmi", middlewares.SetMiddlewareJSON(s.GetMetricBMI)).Methods("GET")
	s.Router.HandleFunc("/imperial/bmi", middlewares.SetMiddlewareJSON(s.GetImperialBMI)).Methods("GET")

	// TEE
	s.Router.HandleFunc("/tee", middlewares.SetMiddlewareJSON(s.GetTotalEnergyExpenditure)).Methods("GET")

	// Swagger Docs:
	opts := middleware.RedocOpts{SpecURL: "../../swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	s.Router.Handle("/docs", sh)
	s.Router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
}
