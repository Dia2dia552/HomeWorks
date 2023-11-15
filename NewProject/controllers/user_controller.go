package controllers

import (
	"NewProject/models"
	"encoding/json"
	"net/http"
)

type UserController struct{}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{ID: 1, Name: "John", Email: "john@example.com"}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(newUser)
	if err != nil {
		return
	}
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}
