package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenValidate(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Token não encontrado",
		})
		c.Abort()
		return
	}

	if len(token) < 16 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Token inválido",
		})
		c.Abort()
		return
	}

	c.Next()

}
