package task_http

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/delivery/http/requests"
	"fun-with-gin/pkg/common"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"net/http"
)

func (h *taskInterActor) UpdateTask(c *gin.Context) {
	ctx := c.Request.Context()
	userMeta, err := common.GetMetaHttpToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	taskId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	validate = validator2.New(validator2.WithRequiredStructEnabled())
	var payload requests.UpdateTaskRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := entity.NewTask(entity.TaskDto{
		UserId:      taskId,
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
	})

	errUpdate := h.taskUC.UpdateOne(ctx, data, &entity.TaskFilter{ID: taskId, UserId: userMeta.Id})
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
