package talkerController

import (
	"main/src/database"
	"main/src/models"
	talkerRepository "main/src/repositories/talker-repository"
	talkerUsecase "main/src/useCases/talker-usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	db, err := database.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	talker := models.Talker{}
	if err = c.BindJSON(&talker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	repository := talkerRepository.NewTalkerRepository(db)
	useCase := talkerUsecase.NewTalkerUseCase(repository)
	if err := useCase.Create(talker); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Talker created",
	})
}
