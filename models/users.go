package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Password     string    `json:"password"`
}
type CreateUser struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
type UpdateUser struct {
	ID           uint   `json:"id" binding:"required"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Password     string `json:"password"`
}
type FindByIdUser struct {
	ID uint `json:"id" binding:"required"`
}
type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SigninUser struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
