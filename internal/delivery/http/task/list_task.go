package task_http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *taskInterActor) ListTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
