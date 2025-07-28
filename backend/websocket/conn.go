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
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID    string
	Email     string
	Conn      *websocket.Conn
	IP        string
	UserAgent string
	ConnTime  time.Time
}

type Message struct {
	To      string `json:"to"`      // recipient UserID
	Message string `json:"message"` // actual message content
}

var clients = make(map[string]*Client) // key: UserID

func (c *Client) Register() {
	clients[c.UserID] = c
}

func (c *Client) Unregister() {
	if c.Conn != nil {
		c.Conn.Close()
	}
	delete(clients, c.UserID)
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
	if err != nil || parsedToken == nil || !parsedToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || claims["id"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID := fmt.Sprintf("%v", claims["id"])
	email := fmt.Sprintf("%v", claims["email"])

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	client := &Client{
		UserID:    userID,
		Email:     email,
		Conn:      conn,
		IP:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		ConnTime:  time.Now(),
	}
	client.Register()
	defer client.Unregister()
	defer conn.Close()

	welcomeMsg, _ := json.Marshal(map[string]interface{}{"message": "Welcome to the WebSocket server!", "userID": userID})
	conn.WriteMessage(websocket.TextMessage, welcomeMsg)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error (user %s): %v", client.UserID, err)
			break
		}

		var incoming Message
		if err := json.Unmarshal(msg, &incoming); err != nil {
			log.Printf("Invalid message format from %s: %v", client.UserID, err)
			continue
		}

		log.Printf("User %s wants to send message to %s: %s", client.UserID, incoming.To, incoming.Message)

		if receiver, ok := clients[incoming.To]; ok && receiver.Conn != nil {
			outgoing, _ := json.Marshal(map[string]interface{}{
				"from":    client.UserID,
				"message": incoming.Message,
			})
			if err := receiver.Conn.WriteMessage(websocket.TextMessage, outgoing); err != nil {
				log.Printf("Error sending message to %s: %v", incoming.To, err)
			}
		} else {
			log.Printf("User %s not connected", incoming.To)
		}
	}
}
