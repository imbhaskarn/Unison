package websocket

import (
	"Unison/helpers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID    string
	Email     string
	ID        string
	Conn      *websocket.Conn
	IP        string
	UserAgent string
	ConnTime  time.Time
}

var clients = make(map[string]*Client)

func (c *Client) Register() {
	clients[c.ID] = c
}

func (c *Client) Unregister() {
	if c.Conn != nil {
		c.Conn.Close()
	}
	delete(clients, c.ID)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections by default; adjust as needed for security
	},
}

func HandleWebSocket(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	log.Println("Authorization header:", authHeader)
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	parsedToken, err := helpers.VerifyToken(authHeader)
	if err != nil || parsedToken == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	if !parsedToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || claims["id"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}
	c.Set("userID", fmt.Sprintf("%v", claims["id"]))
	c.Set("email", claims["email"])

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	client := &Client{
		UserID:    c.GetString("userID"),
		Email:     c.GetString("email"),
		ID:        uuid.New().String(), // Unique UUID
		Conn:      conn,
		IP:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		ConnTime:  time.Now(),
	}
	client.Register()

	print("email:", c.GetString("email"))

	defer client.Unregister()
	defer conn.Close()

	welcomeMsg, _ := json.Marshal(map[string]interface{}{"message": "Welcome to the WebSocket server!", "client": client})

	conn.WriteMessage(websocket.TextMessage, welcomeMsg)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error (conn %s): %v", client.ID, err)
			break
		}
		log.Printf("Received from %s: %s", client.ID, msg)
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Printf("WebSocket write error (conn %s): %v", client.ID, err)
			break
		}
	}
}
