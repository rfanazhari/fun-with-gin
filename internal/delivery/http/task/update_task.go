package task_http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *taskInterActor) UpdateTask(c *gin.Context) {
	ctx := c.Request.Context()
	err := h.taskUC.CreateOne(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
