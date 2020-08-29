package users

import (
	"d-api/databases"
	"d-api/models"
)

func FindAll() []models.User {
	db := databases.GetDb()
	var users []models.User
	db.Find(&users)
	return users
}

func FindById(userId uint) (err error, user models.User) {
	db := databases.GetDb()
	err = db.First(&user, userId).Error
	return err, user
}

func FindByEmail(email string) (err error, user models.User) {
	db := databases.GetDb()
	err = db.Where(&models.User{Email: email}).First(&user).Error
	return err, user
}

func Create(payload models.CreateUser) (err error, user models.User) {
	db := databases.GetDb()
	user = models.User{Email: payload.Email, Name: payload.Name}
	err = db.Create(&user).Error
	return err, user
}

func Update(payload models.UpdateUser) (err error, user models.User) {
	db := databases.GetDb()
	err = db.First(&user, payload.ID).Error
	if err != nil {
		return err, user
	}

	err = db.Model(&user).Update(payload).Error
	return err, user
}

func Delete(userId uint) (err error, user models.User) {
	db := databases.GetDb()
	err = db.Where(&models.User{ID: userId}).Delete(&user).Error
	return err, user
}
