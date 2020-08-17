package controllers

func (s *Server) InitializeRoutes() {
	// Home
	s.Router.HandleFunc("/", s.Home).Methods("GET")
}
