package user_http

import (
	"fun-with-gin/internal/delivery/http/requests"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator2.Validate

func (h *userInterActor) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var payload requests.LoginRequest
	validate = validator2.New(validator2.WithRequiredStructEnabled())

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userUC.LoginUser(ctx, payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})
}
