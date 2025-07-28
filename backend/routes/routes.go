package routers

import (
	"Unison/routes/auth"
	"Unison/routes/document"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	authGroup := r.Group("/auth")

	authGroup.Use(gin.Logger())
	auth.RegisterRoutes(authGroup)
	document.RegisterRoutes(r.Group("/documents"))

	return r
}
