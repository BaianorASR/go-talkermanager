package routes

import (
	talkerController "main/src/controllers/talker-controller"
	"main/src/middleware"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()
	}
}

func talkerRoute(router *gin.RouterGroup) {
	router.Use(Cors())
	r := router.Group("/talker")

	r.GET("/", talkerController.GetAllController)
	r.POST("/", middleware.TokenValidate, middleware.TalkerValidate, talkerController.CreateController)
	r.PUT("/:id", middleware.TokenValidate, middleware.TalkerValidate, talkerController.UpdateController)
	r.GET("/:id", talkerController.GetByIdController)
	r.DELETE("/:id", talkerController.GetByIdController)

}
