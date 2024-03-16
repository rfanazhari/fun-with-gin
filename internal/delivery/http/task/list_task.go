package task_http

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/delivery/http/responses"
	"fun-with-gin/pkg/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *taskInterActor) ListTasks(c *gin.Context) {
	ctx := c.Request.Context()
	userMeta, err := common.GetMetaHttpToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	tasks, err := h.taskUC.ListTask(ctx, &entity.TaskFilter{UserId: userMeta.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": responses.ToListTaskResponse(tasks)})
}
