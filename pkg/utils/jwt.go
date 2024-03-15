package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var (
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	JwtDuration  = os.Getenv("JWT_DURATION")
)

type JwtPayload struct {
	Id        uint   `json:"jti,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
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

func (p *JwtPayload) Valid() error {
	if p.Name == "" {
		return errors.New("name is required")
	}

	if time.Now().Unix() > p.ExpiresAt {
		return errors.New("token has expired")
	}
	return nil
}

func (p *JwtPayload) GetIssuedAt() time.Time {
	return time.Now()
}

func (p *JwtPayload) GetExpiresAt() time.Time {
	return time.Now().Add(time.Hour)
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		token, err := parseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			return
		}

		claims, ok := token.Claims.(*JwtPayload)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token claims"})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

func parseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecretKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
