package talkerController

import (
	"main/src/database"
	talkerRepository "main/src/repositories/talker-repository"
	talkerUsecase "main/src/useCases/talker-usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteController(c *gin.Context) {
	paramID := c.Param("id")
	intID, _ := strconv.Atoi(paramID)

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

	if err := useCase.Delete(intID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    http.StatusOK,
		"message":   "Talker Deleted",
		"talker_id": intID,
	})
}
