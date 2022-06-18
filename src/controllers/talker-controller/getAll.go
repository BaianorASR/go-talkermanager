package talkerController

import (
	"main/src/database"
	talkerRepository "main/src/repositories/talker-repository"
	talkerUsecase "main/src/useCases/talker-usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllController(c *gin.Context) {
	db, err := database.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	repository := talkerRepository.NewTalkerRepository(db)

	useCase := talkerUsecase.NewTalkerUseCasae(repository)

	talkers, err := useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, talkers)
}
