package main

import (
	user_http "fun-with-gin/internal/delivery/http/user"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		panic(err)
	}

	r.Use(utils.LoggerMiddleware())

	user_http.NewUserHandler(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
