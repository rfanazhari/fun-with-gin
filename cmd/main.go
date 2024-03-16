package main

import (
	task_http "fun-with-gin/internal/delivery/http/task"
	user_http "fun-with-gin/internal/delivery/http/user"
	pgsql_task "fun-with-gin/internal/repository/pgsql/task"
	pgsql_user "fun-with-gin/internal/repository/pgsql/user"
	"fun-with-gin/internal/service/redis/login_token"
	"fun-with-gin/pkg/config"
	redisClient "fun-with-gin/pkg/service/redis"
	"fun-with-gin/pkg/utils"
	"fun-with-gin/pkg/validations"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func main() {
	// init DB Postgresql
	db, err := config.NewConnectionPGSQL()
	if err != nil {
		panic(err)
	}

	// init Conn Redis & Call StoreTokenService
	rDB := config.NewRedisConnection()
	newRedisClient := redisClient.NewRedisClient(rDB)
	storeTokenService := login_token.NewRedisLoginToken(newRedisClient)

	// Call Repository
	userRepo := pgsql_user.NewUserPgsqlInit(db)
	taskRepo := pgsql_task.NewTaskPgsqlInit(db)

	// Define Gin
	r := gin.Default()

	// Add Custom Binding
	validations.RegisterBindingValidation()

	r.ForwardedByClientIP = true
	if os.Getenv("PROXY_URL") != "" {
		proxyArray := strings.Split(os.Getenv("PROXY_URL"), ";")
		if len(proxyArray) > 0 {
			err := r.SetTrustedProxies(proxyArray)
			if err != nil {
				panic(err)
			}
		}
	}

	r.Use(utils.LoggerMiddleware())

	user_http.NewUserHandler(r, userRepo, storeTokenService)
	task_http.NewTaskHandler(r, taskRepo, storeTokenService)

	if err := r.Run(os.Getenv("APP_PORT")); err != nil {
		panic(err)
	}
}
