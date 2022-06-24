package talkerRepository

import "errors"

func (r *talkerRepository) Delete(id int) error {
	for i, T := range r.DB.Talkers {
		if T.ID == id {
			println(T.ID, id)

			r.DB.Talkers = append(r.DB.Talkers[:i], r.DB.Talkers[i+1:]...)
			if err := r.DB.Save(); err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("Talker not found for delete")
}
