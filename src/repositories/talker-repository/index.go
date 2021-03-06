package talkerRepository

import (
	"main/src/database"
	"main/src/models"
)

type talkerRepository struct {
	DB *database.Database
}

type UserRepository interface {
	GetAll() ([]models.Talker, error)
	GetByID(id int) (models.Talker, error)
	Create(talker models.Talker) error
	Update(talker models.Talker) error
	Delete(id int) error
	Search(query string) []models.Talker
}

func NewTalkerRepository(db *database.Database) UserRepository {
	return &talkerRepository{db}
}
