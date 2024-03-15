package user_http

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/delivery/http/requests"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"net/http"
)

func (h *userInterActor) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	validate = validator2.New(validator2.WithRequiredStructEnabled())
	var payload requests.UpdateUserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := entity.NewUser(entity.UserDto{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	})

	errUpdate := h.userUC.UpdateOne(ctx, data, &entity.UserFilter{ID: &userId})
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Update User"})
}
