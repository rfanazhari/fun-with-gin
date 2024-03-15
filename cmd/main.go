package main

import (
	task_http "fun-with-gin/internal/delivery/http/task"
	user_http "fun-with-gin/internal/delivery/http/user"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	"os"
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
	task_http.NewTaskHandler(r)

	if err := r.Run(os.Getenv("APP_PORT")); err != nil {
		panic(err)
	}
}
