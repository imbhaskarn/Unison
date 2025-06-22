package routers

import (
	"Unison/db"
	"Unison/routes/auth"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Unison API",
		})
	})
		r.GET("/db", func(c *gin.Context) {
		db.DB.Ping() // Check if the database connection is alive
		if err := db.DB.Ping(); err != nil {
			c.JSON(500, gin.H{
				"error": "Database connection failed",
			})
			return
		}
		_, _ = db.DB.Exec("SELECT 1") // Simple query to ensure the database is operational
		log.Fatal("Database connection is healthy")
		c.JSON(200, gin.H{
			"message": "Welcome to Unison API",
		})
	})
	// Grouping auth routes under /auth
	authGroup := r.Group("/auth")
	auth.RegisterRoutes(authGroup)

	return r
}
