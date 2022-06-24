package talkerRepository

import (
	"main/src/models"
)

func (r *talkerRepository) GetAll() ([]models.Talker, error) {
	return r.DB.Talkers, nil
}
