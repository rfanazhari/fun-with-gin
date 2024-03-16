package task_http

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/delivery/http/responses"
	"fun-with-gin/pkg/common"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *taskInterActor) FindById(c *gin.Context) {
	ctx := c.Request.Context()
	userMeta, err := common.GetMetaHttpToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	taskId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.taskUC.FindOne(ctx, &entity.TaskFilter{ID: taskId, UserId: userMeta.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": responses.ToTaskResponse(task)})
}
