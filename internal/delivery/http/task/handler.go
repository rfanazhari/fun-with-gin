package task_http

import (
	"context"
	"fun-with-gin/domain/usecase"
	usecase_user "fun-with-gin/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type taskInterActor struct {
	userUC usecase.UserUseCase
}

func NewTaskHandler(route *gin.Engine) {
	ctx := context.TODO()

	userUseCase := usecase_user.NewUseCaseUser(ctx)
	handler := &taskInterActor{
		userUC: userUseCase,
	}

	userRoutes := route.Group("/tasks")
	{
		userRoutes.POST("/create", handler.CreateTask)
		userRoutes.GET("/:id", handler.FindById)
		userRoutes.PUT("/:id", handler.UpdateTask)
		userRoutes.GET("", handler.ListTasks)
	}
}
