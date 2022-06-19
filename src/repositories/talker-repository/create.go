package talkerRepository

import (
	"errors"
	"main/src/models"
)

func (r *talkerRepository) Create(talker models.Talker) error {
	var talkers = r.DB.Talkers

	for _, t := range r.DB.Talkers {
		if t.Name == talker.Name {
			return errors.New("Talker already exists")
		}
	}

	talker.ID = talkers[len(talkers)-1].ID + 1
	r.DB.Talkers = append(talkers, talker)

	if err := r.DB.Save(); err != nil {
		return err
	}

	return nil
}
