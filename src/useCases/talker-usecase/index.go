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
	// Create(talker *database.Database) error
	// Update(talker *database.Database) error
	// Delete(id int) error
}

func NewTalkerUseCasae(r talkerRepository.UserRepository) UseCase {
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
