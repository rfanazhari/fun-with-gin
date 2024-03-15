package user_http

import (
	"context"
	"fun-with-gin/domain/usecase"
	usecase_user "fun-with-gin/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type userInterActor struct {
	userUC usecase.UserUseCase
}

func NewUserHandler(route *gin.Engine) {
	ctx := context.TODO()

	userUseCase := usecase_user.NewUseCaseUser(ctx)
	handler := &userInterActor{
		userUC: userUseCase,
	}

	route.POST("/login", handler.Login)

	userRoutes := route.Group("/users")
	{
		userRoutes.POST("/create", handler.CreateUser)
		userRoutes.GET("/:id", handler.FindById)
		userRoutes.PUT("/:id", handler.UpdateUser)
		userRoutes.GET("", handler.ListUsers)
	}
}
