package responses

import (
	"fun-with-gin/domain/entity"
	"time"
)

type TaskResponse struct {
	ID          uint      `json:"ID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ToTaskResponse(task *entity.Task) *TaskResponse {
	return &TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func ToListTaskResponse(tasks []*entity.Task) []*TaskResponse {
	var result []*TaskResponse
	for _, task := range tasks {
		result = append(result, ToTaskResponse(task))
	}
	return result
}
