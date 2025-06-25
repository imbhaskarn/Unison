package main

import (
	"Unison/db"
	"Unison/routes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)
	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	router := routers.SetupRouter()

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}
			log.Printf("Received: %s", msg)
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("WebSocket write error:", err)
				break
			}
		}
	})

	log.Println("server is starting on: http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
