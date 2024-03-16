package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func ConvertBool(env string) bool {
	v, _ := strconv.ParseBool(os.Getenv(env))
	return v
}

func DefaultValueDuration(env string, defaultValue string) time.Duration {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		panic(err)
	}
	return d
}

func DefaultValueBool(env string, defaultValue bool) bool {
	val, err := strconv.ParseBool(os.Getenv(env))
	if err != nil {
		val = false
	}

	if val == false {
		return defaultValue
	}
	return val
}

func GenerateID() int64 {
	return time.Now().Unix()
}

func StringToUint(s string) (uint, error) {
	uintVal, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(uintVal), nil
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		fmt.Printf("[%s] %s %s %s %d\n", time.Now().Format("2006-01-02 15:04:05"), c.Request.Method, c.Request.URL.Path, c.Request.Proto, c.Writer.Status())

		fmt.Printf("Duration: %v\n", time.Since(start))
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidatePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func UintToString(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}
