package talkerController

import (
	"main/src/database"
	"main/src/models"
	talkerRepository "main/src/repositories/talker-repository"
	talkerUsecase "main/src/useCases/talker-usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {
	paramID := c.Param("id")
	intId, err := strconv.Atoi(paramID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID",
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

	talker := models.Talker{}
	talker.ID = intId

	if err = c.BindJSON(&talker); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	repository := talkerRepository.NewTalkerRepository(db)
	useCase := talkerUsecase.NewTalkerUseCase(repository)
	if err := useCase.Update(talker); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Talker updated",
		"talker":  talker,
	})
}
