package routes

import (
	talkerController "main/src/controllers/talker-controller"
	"main/src/middleware"

	"github.com/gin-gonic/gin"
)

func talkerRoute(router *gin.RouterGroup) {
	r := router.Group("/talker")
	{
		r.GET("/", talkerController.GetAllController)
		r.POST("/", middleware.TokenValidate, middleware.TalkerValidate, talkerController.CreateController)
		r.GET("/:id", talkerController.GetByIdController)
		r.PUT("/:id", talkerController.GetByIdController)
		r.DELETE("/:id", talkerController.GetByIdController)
	}
}
