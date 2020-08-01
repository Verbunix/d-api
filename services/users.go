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

func FindUserById(userId uint) (err error, user models.User) {
	db := databases.GetDb()
	err = db.First(&user, userId).Error
	return err, user
}

func FindUserByEmail(email string) (err error, user models.User) {
	db := databases.GetDb()
	err = db.Where(&models.User{Email: email}).First(&user).Error
	return err, user
}

func CreateUser(payload models.CreateUser) (err error, user models.User) {
	db := databases.GetDb()
	user = models.User{Email: payload.Email, Name: payload.Name}
	err = db.Create(&user).Error
	return err, user
}

func UpdateUser(payload models.UpdateUser) (err error, user models.User) {
	db := databases.GetDb()
	err = db.First(&user, payload.ID).Error
	if err != nil {
		return err, user
	}

	err = db.Model(&user).Update(payload).Error
	return err, user
}
