package requests

import _ "fun-with-gin/pkg/validations"

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required,max=255"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required,max=50"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"omitempty,maxLengthWhenNotEmpty"`
	Description string `json:"description" binding:"omitempty"`
	Status      string `json:"status" binding:"omitempty,maxLengthWhenNotEmpty"`
}
