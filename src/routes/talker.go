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

	r.GET("/:id", talkerController.GetByIdController)
	r.GET("/", talkerController.GetAllController)
	{
		r.Use(middleware.TokenValidate)
		r.GET("/search", talkerController.SearchController)
		r.DELETE("/:id", talkerController.DeleteController)
		{
			r.Use(middleware.TalkerValidate)
			r.POST("/", talkerController.CreateController)
			r.PUT("/:id", talkerController.UpdateController)
		}
	}

}
