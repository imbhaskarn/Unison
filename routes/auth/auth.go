package auth

import (
	authHandlers "Unison/handlers/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", authHandlers.Register)
}
