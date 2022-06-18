package routes

import (
	"main/src/controllers"
	"main/src/middleware"

	"github.com/gin-gonic/gin"
)

func loginRoute(superRoute *gin.RouterGroup) {

	r := superRoute.Group("/login")
	{
		r.POST("/", middleware.LoginValidation, controllers.LoginController)
	}
}
