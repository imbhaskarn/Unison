package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"Unison/db"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		log.Println("Received a request on /")
		c.String(200, "Hello, World!")
	})
	log.Println("Server is starting on :8080")
	log.Fatal(router.Run(":8080"))
}
