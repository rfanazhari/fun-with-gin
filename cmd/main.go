package main

import (
	task_http "fun-with-gin/internal/delivery/http/task"
	user_http "fun-with-gin/internal/delivery/http/user"
	pgsql_task "fun-with-gin/internal/repository/pgsql/task"
	pgsql_user "fun-with-gin/internal/repository/pgsql/user"
	"fun-with-gin/pkg/config"
	"fun-with-gin/pkg/utils"
	"fun-with-gin/pkg/validations"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func main() {
	db, err := config.NewConnectionPGSQL()
	if err != nil {
		panic(err)
	}

	userRepo := pgsql_user.NewUserPgsqlInit(db)
	taskRepo := pgsql_task.NewTaskPgsqlInit(db)

	r := gin.Default()

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

	user_http.NewUserHandler(r, userRepo)
	task_http.NewTaskHandler(r, taskRepo)

	if err := r.Run(os.Getenv("APP_PORT")); err != nil {
		panic(err)
	}
}
