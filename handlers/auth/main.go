package authHandlers

import (
	"Unison/db"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login handles user login requests.
func Login(c *gin.Context) {
	var resBody map[string]any
	if err := c.ShouldBindJSON(&resBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
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

	findUser := db.DB.QueryRow("Select id, \"hashedPassword\" from users where email = $1", email)

	type User struct {
		id             int
		hashedPassword string
	}

	var user User
	if err := findUser.Scan(&user.id, &user.hashedPassword); err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println("Error querying database:", err)
			c.JSON(500, gin.H{"error": "Database error1"})
			return
		}
	}
	if user.id == 0 {
		log.Println("User not found:", email)
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.hashedPassword), []byte(password.(string)))

	if err != nil {
		log.Println("Invalid password:", err)
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	data, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    user.id,
		"exp":   jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}).SignedString([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Println("Error generating JWT:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	log.Println("User logged in successfully:", email)
	c.JSON(200, gin.H{"message": "Login successful", "token": data})
}

// Register handles user registration requests.
func Register(c *gin.Context) {
	var resBody map[string]any
	if err := c.ShouldBindJSON(&resBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
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

	findUser := db.DB.QueryRow("Select id from users where email = $1", email)
	var userID int
	if err := findUser.Scan(&userID); err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println("Error querying database:", err)
			c.JSON(500, gin.H{"error": "Database error1"})
			return
		}
	}
	log.Println("User ID found:", userID)
	if userID != 0 {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	var err error
	hash, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	log.Println("Password hashed successfully", string(hash))
	_, err = db.DB.Exec("INSERT INTO users (email, \"hashedPassword\") VALUES ($1, $2)", email.(string), string(hash))

	if err != nil {
		log.Println("Error inserting user into database:", err)
		c.JSON(500, gin.H{"error": "Database error2", "details": err.Error()})
		return
	}
	log.Println("User registered successfully:", email)

	data, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	}).SignedString([]byte(os.Getenv("SESSION_SECRET")))

	if err != nil {
		log.Println("Error generating JWT:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	log.Println("JWT generated successfully:", data)

	c.JSON(201, gin.H{"message": "User registered successfully", "token": data})
}

// RefreshToken handles token refresh requests.
func RefreshToken(c *gin.Context) {
	// TODO: Implement token refresh logic
	c.JSON(200, gin.H{"message": "Refresh token handler"})
}
