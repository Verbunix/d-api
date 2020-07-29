package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	AccessToken  string    `json:"access_token";sql:"-"`
	RefreshToken string    `json:"refresh_token";sql:"-"`
	Password     string    `json:"password";sql:"-"`
}
type CreateUser struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
type UpdateUser struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type FindByIdUser struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name"`
}
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
