package auth

import (
	authHandlers "Unison/handlers/auth"
	"Unison/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/signup", authHandlers.Register)
	router.POST("/login", authHandlers.Login)
	router.GET("/user", middlewares.AuthRequired(), authHandlers.UserHandler)
}
