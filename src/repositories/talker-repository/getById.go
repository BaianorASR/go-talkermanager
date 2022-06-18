package talkerRepository

import (
	"errors"
	"main/src/models"
)

func (r *talkerRepository) GetByID(id int) (models.Talker, error) {
	for _, talker := range r.DB.Talkers {
		if talker.ID == id {
			return talker, nil
		}
	}

	return models.Talker{}, errors.New("Talker not found")
}
