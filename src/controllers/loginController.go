package controllers

import "github.com/gin-gonic/gin"

func LoginController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Baianor",
	})
}
