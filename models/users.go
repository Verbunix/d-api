package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
type CreateUser struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
type UpdateUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
