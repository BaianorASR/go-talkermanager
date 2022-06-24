package talkerController

import (
	"main/src/database"
	talkerRepository "main/src/repositories/talker-repository"
	talkerUsecase "main/src/useCases/talker-usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchController(c *gin.Context) {
	query := c.Query("q")
	println(query)

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

	talkers := useCase.Search(query)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": talkers,
	})
}
