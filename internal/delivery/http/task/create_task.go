package task_http

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/delivery/http/requests"
	"fun-with-gin/pkg/common"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"net/http"
)

func (h *taskInterActor) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	var payload requests.CreateTaskRequest
	validate = validator2.New(validator2.WithRequiredStructEnabled())

	userMeta, err := common.GetMetaHttpToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := entity.NewTask(entity.TaskDto{
		UserId:      userMeta.Id,
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
	})

	errUc := h.taskUC.CreateOne(ctx, task)
	if errUc != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUc.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
