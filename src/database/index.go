package database

import (
	"encoding/json"
	"io/ioutil"
	"main/src/models"
)

type Database struct {
	Talkers []models.Talker `json:"talkers"`
}

func Save(t *[]models.Talker) error {
	file, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("src/database/talker.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Get() (*Database, error) {
	file, err := ioutil.ReadFile("src/database/talker.json")
	if err != nil {
		return nil, err
	}

	var db = Database{}
	err = json.Unmarshal([]byte(file), &db.Talkers)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
