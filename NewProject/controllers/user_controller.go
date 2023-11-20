package controllers

import (
	"NewProject/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

type UserController struct {
	db *sql.DB
}

func (uc *UserController) SetDB(database *sql.DB) {
	uc.db = database
}

func (uc *UserController) AddUsers(w http.ResponseWriter, r *http.Request) error {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		return err
	}

	_, err = uc.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", newUser.Name, newUser.Email)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) ([]models.User, error) {
	rows, err := uc.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
		}
	}()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
