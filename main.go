package main

import (
	"Unison/db"
	"Unison/routes"
	"github.com/gin-gonic/gin"
	"Unison/websocket"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)
	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	router := routers.SetupRouter()



	router.GET("/ws", websocket.HandleWebSocket)

	log.Println("server is starting on: http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
