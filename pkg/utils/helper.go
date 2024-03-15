package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func GenerateID() uint64 {
	randomNumber := rand.Intn(900000) + 100000
	currentTime := time.Now().UnixNano()
	randomID := fmt.Sprintf("%d%d", randomNumber, currentTime)

	var randomIDUint uint64
	fmt.Sscanf(randomID, "%d", &randomIDUint)

	return randomIDUint
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
