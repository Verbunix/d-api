package controllers

import (
	"dating-api/models"
	"dating-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var createUser models.CreateUser

	if err := c.ShouldBindJSON(&signinUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	createUser.Email = signinUser.Email
	createUser.Name = signinUser.Name

	err, u := services.CreateUser(createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err, token := services.CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err, uUpdated := services.UpdateUser(models.UpdateUser{ID: u.ID, Password: token})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, uUpdated)
}
