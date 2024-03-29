package user_http

import (
	"fun-with-gin/internal/delivery/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *userInterActor) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.userUC.ListUser(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": responses.ToListUserResponses(users)})
}
