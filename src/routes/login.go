package routes

import (
	"main/src/controllers"

	"github.com/gin-gonic/gin"
)

func loginRoute(superRoute *gin.RouterGroup) {
	superRoute.GET("/login", controllers.LoginController)
	superRoute.POST("/login", controllers.LoginController)
}
