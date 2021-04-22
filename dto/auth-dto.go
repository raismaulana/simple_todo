package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
}
