package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaim struct {
}

func GenerateJWTToken(secretKey string, name string, id int) (string, error) {
	userClaims := jwt.MapClaims{
		"username": name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// Sign and get the complete encoded token as a string using the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
