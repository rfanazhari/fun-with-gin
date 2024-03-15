package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	JWT_DURATION   = os.Getenv("JWT_DURATION")
)

type JwtPayload struct {
	Id    uint
	Name  string
	Email string
}

func GenerateToken(payload JwtPayload) (string, error) {
	jwtDurationString := os.Getenv("JWT_DURATION")
	jwtDuration, err := time.ParseDuration(jwtDurationString)
	if err != nil {
		panic("Invalid JWT_DURATION value")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    payload.Id,
		"name":  payload.Name,
		"email": payload.Email,
		"exp":   time.Now().Add(jwtDuration).Unix(), // Token expires in 24 hours
	})
	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
