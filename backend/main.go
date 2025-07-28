package main

import (
	"Unison/db"
	"Unison/routes"
	"Unison/websocket"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)
	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	router := routers.SetupRouter()

router.Use(cors.New(cors.Config{
  AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
  AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
  AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
  ExposeHeaders:    []string{"Content-Length"},
  AllowCredentials: true,
  MaxAge: 12 * time.Hour,
}))


	router.GET("/ws", websocket.HandleWebSocket)

	log.Println("server is starting on: http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
