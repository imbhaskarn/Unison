package authHandlers

import (
	"Unison/db"
	"log"

	"github.com/gin-gonic/gin"
)

// Login handles user login requests.
func Login(c *gin.Context) {
	// TODO: Implement login logic
	c.String(200, "Login handler")
}

// Register handles user registration requests.
func Register(c *gin.Context) {
	var resBody map[string]interface{}
	if err := c.ShouldBindJSON(&resBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	email := resBody["email"]
	if email == nil {
		c.JSON(400, gin.H{"error": "Email is required"})
		return
	}
	password := resBody["password"]
	if password == nil {
		c.JSON(400, gin.H{"error": "Password is required"})
		return
	}

	findUser := db.DB.QueryRow("Select id from users where email = ?", email)
	var userID int
	if err := findUser.Scan(&userID); err != nil {
		if err.Error() != "sql: no rows in result set" {
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}
	}
	log.Println("User ID found:", userID)
	if userID != 0 {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

// RefreshToken handles token refresh requests.
func RefreshToken(c *gin.Context) {
	// TODO: Implement token refresh logic
	c.JSON(200, gin.H{"message": "Refresh token handler"})
}
