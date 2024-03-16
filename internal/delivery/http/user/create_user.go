package user_http

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/delivery/http/requests"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"net/http"
)

func (h *userInterActor) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	validate = validator2.New(validator2.WithRequiredStructEnabled())
	var payload requests.CreateUserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, errPassword := utils.HashPassword(payload.Password)
	if errPassword != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errPassword.Error()})
		return
	}
	user := entity.NewUser(entity.UserDto{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: password,
	})

	err := h.userUC.CreateOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
