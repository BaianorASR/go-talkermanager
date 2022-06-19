package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func LoginController(c *gin.Context) {
	id, err := gonanoid.New(16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"source":  c.Request.URL.Path,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  id,
		"source": c.Request.URL.Path,
	})
}
