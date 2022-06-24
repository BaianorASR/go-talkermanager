package talkerRepository

import (
	"errors"
	"main/src/models"
)

func (r *talkerRepository) Update(talker models.Talker) error {
	for index, T := range r.DB.Talkers {
		if T.ID == talker.ID {
			r.DB.Talkers[index] = talker
			if err := r.DB.Save(); err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("Talker not found")
}
