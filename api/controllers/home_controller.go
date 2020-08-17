package controllers

import (
	"encoding/json"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Welcome to Nutriport API")
}
