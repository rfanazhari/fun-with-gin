package task_http

import (
	"context"
	"fun-with-gin/domain/repository"
	"fun-with-gin/domain/service"
	"fun-with-gin/domain/usecase"
	usecase_task "fun-with-gin/internal/usecase/task"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
)

type taskInterActor struct {
	taskUC usecase.TaskUseCase
}

var validate *validator2.Validate

func NewTaskHandler(route *gin.Engine, taskRepo repository.TaskRepository, tokenSvc service.LoginTokenService) {
	ctx := context.TODO()

	userUseCase := usecase_task.NewUseCaseTask(taskRepo)
	handler := &taskInterActor{
		taskUC: userUseCase,
	}

	route.Use(utils.JwtMiddleware(ctx, tokenSvc))

	userRoutes := route.Group("/tasks")
	{
		userRoutes.POST("/create", handler.CreateTask)
		userRoutes.GET("/:id", handler.FindById)
		userRoutes.PUT("/:id", handler.UpdateTask)
		userRoutes.GET("", handler.ListTasks)
	}
}
