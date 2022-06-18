package talkerRepository

import (
	"errors"
	"main/src/models"
)

func (r *talkerRepository) GetAll() ([]models.Talker, error) {
	if len(r.DB.Talkers) == 0 {
		return nil, errors.New("Talkers not found")
	}

	return r.DB.Talkers, nil
}
