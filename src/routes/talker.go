package routes

import (
	talkerController "main/src/controllers/talker-controller"

	"github.com/gin-gonic/gin"
)

func talkerRoute(router *gin.RouterGroup) {
	r := router.Group("/talker")
	{
		r.GET("/", talkerController.GetAllController)
		r.POST("/", talkerController.GetByIdController)
		r.GET("/:id", talkerController.GetByIdController)
		r.PUT("/:id", talkerController.GetByIdController)
		r.DELETE("/:id", talkerController.GetByIdController)
	}
}
