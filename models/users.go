package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"current_time"`
}
type CreateUser struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
type UpdateUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
