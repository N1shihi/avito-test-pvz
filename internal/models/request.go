package models

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64,alphanum"`
	Role     string `json:"role" binding:"required,oneof=employee moderator"`
}
