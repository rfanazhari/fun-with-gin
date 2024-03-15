package task_http

import (
	"context"
	"fun-with-gin/domain/usecase"
	usecase_task "fun-with-gin/internal/usecase/task"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type taskInterActor struct {
	userUC usecase.TaskUseCase
}

func NewTaskHandler(route *gin.Engine) {
	ctx := context.TODO()

	userUseCase := usecase_task.NewUseCaseTask(ctx)
	handler := &taskInterActor{
		userUC: userUseCase,
	}

	route.Use(utils.JwtMiddleware())

	userRoutes := route.Group("/tasks")
	{
		userRoutes.POST("/create", handler.CreateTask)
		userRoutes.GET("/:id", handler.FindById)
		userRoutes.PUT("/:id", handler.UpdateTask)
		userRoutes.GET("", handler.ListTasks)
	}
}
