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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid id",
		})
		return
	}

	db, err := database.Get()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	repository := talkerRepository.NewTalkerRepository(db)

	useCase := talkerUsecase.NewTalkerUseCase(repository)

	talker, err := useCase.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, talker)
}
