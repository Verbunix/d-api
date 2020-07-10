package services

import (
	"dating-api/databases"
	"dating-api/models"
)

func FindUsers() []models.User {
	db := databases.GetDb()
	var users []models.User
	db.Find(&users)
	return users
}

func CreateUser(payload models.CreateUser) (err error, user models.User) {
	db := databases.GetDb()
	user = models.User{Email: payload.Email, Name: payload.Name}
	err = db.Create(&user).Error
	return err, user
}
