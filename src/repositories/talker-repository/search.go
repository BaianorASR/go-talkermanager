package talkerRepository

import (
	"main/src/models"
	"strings"
)

func (r *talkerRepository) Search(query string) []models.Talker {
	var talkers = make([]models.Talker, 0)

	if query == "" {
		return talkers
	}

	for _, T := range r.DB.Talkers {
		if strings.Contains(strings.ToLower(T.Name), strings.ToLower(query)) {
			talkers = append(talkers, T)
		}
	}

	return talkers
}
