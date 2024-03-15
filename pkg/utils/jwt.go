package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	JwtDuration  = os.Getenv("JWT_DURATION")
)

type JwtPayload struct {
	Id    uint
	Name  string
	Email string
}

func GenerateToken(payload JwtPayload) (string, error) {
	jwtDuration, err := time.ParseDuration(JwtDuration)
	if err != nil {
		panic("Invalid JWT_DURATION value")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    payload.Id,
		"name":  payload.Name,
		"email": payload.Email,
		"exp":   time.Now().Add(jwtDuration).Unix(), // Token expires in 24 hours
	})
	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
