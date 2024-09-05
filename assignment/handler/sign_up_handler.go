package handler

import (
	"encoding/json"
	"net/http"

	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/service"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var payload dto.SignUpPayload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// Validate request body
	if err := payload.Validate(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	response, status, err := service.SignUp(payload.Email, payload.Password)
	if status != http.StatusCreated {
		http.Error(w, "", status)
		return
	} else if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SignUpHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		signUp(w, r)
		return
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
