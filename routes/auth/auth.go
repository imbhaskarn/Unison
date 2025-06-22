package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/login", func(c *gin.Context) {
		// Sample login handler
		c.JSON(200, gin.H{
			"message": "Login endpoint hit for get method test",
		})
	})
	rg.POST("/login", func(c *gin.Context) {
		// Sample login handler
		c.JSON(200, gin.H{
			"message": "Login endpoint hit",
		})
	})
}