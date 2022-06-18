package talkerController

import (
	"main/src/database"
	talkerRepository "main/src/repositories/talker-repository"
	talkerUsecase "main/src/useCases/talker-usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetByIdController(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		return
	}

	db, err := database.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	repository := talkerRepository.NewTalkerRepository(db)

	useCase := talkerUsecase.NewTalkerUseCasae(repository)

	talker, err := useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, talker)
}
