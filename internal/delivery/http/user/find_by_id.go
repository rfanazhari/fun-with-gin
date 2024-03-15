package user_http

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *userInterActor) FindById(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filter := &entity.UserFilter{ID: userId}
	user, err := h.userUC.FindOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": user})
}
