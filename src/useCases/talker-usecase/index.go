package talkerUsecase

import (
	"main/src/models"
	talkerRepository "main/src/repositories/talker-repository"
)

type talkerUseCase struct {
	repository talkerRepository.UserRepository
}

type UseCase interface {
	GetAll() ([]models.Talker, error)
	GetByID(id int) (models.Talker, error)
	Create(talker models.Talker) error
	// Update(talker *database.Database) error
	// Delete(id int) error
}

func NewTalkerUseCase(r talkerRepository.UserRepository) UseCase {
	return talkerUseCase{
		repository: r,
	}
}

func (t talkerUseCase) GetAll() ([]models.Talker, error) {
	return t.repository.GetAll()
}

func (t talkerUseCase) GetByID(id int) (models.Talker, error) {
	return t.repository.GetByID(id)
}

func (t talkerUseCase) Create(talker models.Talker) error {
	return t.repository.Create(talker)
}
