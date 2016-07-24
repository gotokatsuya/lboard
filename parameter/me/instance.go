package me

import (
	"encoding/json"

	usermodel "github.com/gotokatsuya/lboard/model/user"
)

type instance struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Convert ...
func (i *instance) Convert(user *usermodel.Entity) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, i); err != nil {
		return err
	}
	return nil
}
