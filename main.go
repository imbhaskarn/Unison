package main

import (
	"Unison/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		log.Println("Received a request on /")
		c.JSON(200, gin.H{
			"message": "Welcome to Unison!",
		})
	})
	log.Println("server is starting on: http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
