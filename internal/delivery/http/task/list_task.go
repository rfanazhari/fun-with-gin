package task_http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *taskInterActor) ListTasks(c *gin.Context) {
	ctx := c.Request.Context()
	tasks, err := h.taskUC.ListTask(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
