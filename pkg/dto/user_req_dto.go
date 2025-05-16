package dto

type RegisterUserRequest struct{
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUserRequest struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}