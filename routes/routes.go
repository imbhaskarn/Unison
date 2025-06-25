package routers

import (
	"Unison/routes/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"}) 
	authGroup := r.Group("/auth")

	authGroup.Use(gin.Logger())
	auth.RegisterRoutes(authGroup)

	return r
}
