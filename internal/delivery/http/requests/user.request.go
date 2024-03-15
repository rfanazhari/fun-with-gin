package requests

type LoginRequest struct {
	Email    string `json:"email" binding:"required,string"`
	Password string `json:"password" binding:"required,email"`
}
