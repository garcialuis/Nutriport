package controllers

import (
	"encoding/json"
	"net/http"
)

// Home handler:
// swagger:route GET / home Home
// Responses:
//	200: description: OK - Welcomes to Nutriport API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Welcome to Nutriport API")
}
