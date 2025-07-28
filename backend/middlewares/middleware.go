// middlewares/auth.go
package middlewares

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized: No token provided"})
			return
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized: Invalid token format"})
			return
		}

		parsedToken, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(os.Getenv("SESSION_SECRET")), nil
		})

		if err != nil {
			fmt.Println("JWT parse error:", err)
		}

		if err != nil || !parsedToken.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized: Invalid token"})
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || claims["id"] == nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized: Invalid token claims"})
			return
		}

		c.Set("userID", fmt.Sprintf("%v", claims["id"]))
		c.Set("email", claims["email"])

		c.Next()
	}
}
