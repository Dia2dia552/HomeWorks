package views

import (
	"NewProject/models"
	"encoding/json"
)

func JsonUser(u *models.User) ([]byte, error) {
	serialized, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return serialized, nil
}
