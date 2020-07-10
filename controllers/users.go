package controllers

import (
	"dating-api/models"
	"dating-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController struct{}

func (u UsersController) FindUsers(c *gin.Context) {
	findUsers := services.FindUsers()
	c.JSON(http.StatusOK, findUsers)
}

func (u UsersController) CreateUsers(c *gin.Context) {
	var input models.CreateUser

	_ = c.BindJSON(&input)
	err, createUser := services.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, createUser)
}
