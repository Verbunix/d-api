package controllers

import (
	"d-api/models"
	"d-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (ac AuthController) Login(c *gin.Context) {
	var u models.LoginUser
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	err, user := services.FindUserByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//get sha256 from request password
	hash := services.CreateShaHash(u.Password)

	//compare the user from the request, with user from db:
	if user.Email != u.Email || user.Password != hash {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	err, token := services.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func (ac AuthController) SignIn(c *gin.Context) {
	var signinUser models.SigninUser

	if err := c.ShouldBindJSON(&signinUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err, u := services.CreateUser(models.CreateUser{Email: signinUser.Email, Name: signinUser.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//get sha256 from request password
	hash := services.CreateShaHash(signinUser.Password)

	//update user and save password hash
	err, uUpdated := services.UpdateUser(models.UpdateUser{ID: u.ID, Password: hash})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, uUpdated)
}
