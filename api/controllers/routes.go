package controllers

import "github.com/garcialuis/Nutriport/api/middlewares"

func (s *Server) InitializeRoutes() {
	// Home
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}
