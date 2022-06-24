package talkerRepository

import (
	"errors"
	"main/src/models"
)

func (r *talkerRepository) Create(talker models.Talker) error {
	for _, t := range r.DB.Talkers {
		if t.Name == talker.Name {
			return errors.New("Talker already exists")
		}
	}

	if len(r.DB.Talkers) == 0 {
		talker.ID = 1
	} else {
		talker.ID = r.DB.Talkers[len(r.DB.Talkers)-1].ID + 1
	}

	r.DB.Talkers = append(r.DB.Talkers, talker)

	if err := r.DB.Save(); err != nil {
		return err
	}

	return nil
}
