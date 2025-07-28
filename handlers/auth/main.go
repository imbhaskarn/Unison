package authHandlers

import (
	"Unison/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

func Login(c *gin.Context) {
	var body map[string]any
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid request body"})
		return
	}
	email := body["email"]
	if email == nil {
		c.JSON(400, gin.H{"success": false, "error": "Email is required"})
		return
	}
	password := body["password"]
	if password == nil {
		c.JSON(400, gin.H{"success": false, "error": "Password is required"})
		return
	}

	findUser := db.DB.QueryRow("SELECT id, \"hashedPassword\" FROM users WHERE email = $1", email)

	var user struct {
		id             int
		hashedPassword string
	}
	if err := findUser.Scan(&user.id, &user.hashedPassword); err != nil {
		log.Println("Database error:", err)
		c.JSON(401, gin.H{"success": false, "error": "Invalid email or password"})
		return
	}
	if user.id == 0 {
		c.JSON(401, gin.H{"success": false, "error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.hashedPassword), []byte(password.(string)))
	if err != nil {
		c.JSON(401, gin.H{"success": false, "error": "Invalid email or password"})
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    user.id,
		"exp":   jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}).SignedString([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Println("JWT generation error:", err)
		c.JSON(500, gin.H{"success": false, "error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Login successful",
		"data": gin.H{
			"accessToken": token,
			"email":       email,
		},
	})
}

func Register(c *gin.Context) {
	var body map[string]any
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid request body"})
		return
	}
	email := body["email"]
	password := body["password"]
	if email == nil || password == nil {
		c.JSON(400, gin.H{"success": false, "error": "Email and password are required"})
		return
	}

	findUser := db.DB.QueryRow("SELECT id, email FROM users WHERE email = $1", email)

	var user struct {
		id    int
		email string
	}
	if err := findUser.Scan(&user.id, &user.email); err == nil && user.id != 0 {
		c.JSON(400, gin.H{"success": false, "error": "User already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Error while hashing password"})
		return
	}

	_, err = db.DB.Exec("INSERT INTO users (email, \"hashedPassword\") VALUES ($1, $2)", email.(string), string(hash))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Error inserting user into database"})
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    user.id,
		"exp":   jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}).SignedString([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "JWT generation failed"})
		return
	}

	c.JSON(201, gin.H{
		"success": true,
		"message": "User registered successfully",
		"data": gin.H{
			"email":       email,
			"accessToken": token,
		},
	})
}

func UserHandler(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, gin.H{"success": false, "error": "Unauthorized"})
		return
	}

	var email string
	err := db.DB.QueryRow("SELECT email FROM users WHERE id = $1", userID).Scan(&email)
	if err != nil {
		c.JSON(404, gin.H{"success": false, "error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User fetched successfully",
		"data": gin.H{
			"userID": userID,
			"email":  email,
		},
	})
}
