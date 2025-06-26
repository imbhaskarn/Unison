package helpers

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenString string) (*jwt.Token, error) {

	token := strings.TrimPrefix(tokenString, "Bearer ")

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(os.Getenv("SESSION_SECRET")), nil
	})


	

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}
