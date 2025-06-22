package main

import (
	"Unison/db"
	"Unison/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	router := routers.SetupRouter()
	router.Use(gin.Recovery())
	
	log.Println("server is starting on: http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
