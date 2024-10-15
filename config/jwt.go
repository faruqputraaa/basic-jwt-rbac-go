package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret")

// Generate JWT Token
func GenerateJWT(userID uint, role string) (string, error) {
	claims := &jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	// Log token untuk memastikan token benar
	fmt.Println("Generated Token:", signedToken)
	return signedToken, nil
}
