package controllers

import (
	"dating-api/models"
	"dating-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UsersController struct{}

func (u UsersController) FindUsers(c *gin.Context) {
	findUsers := services.FindUsers()
	c.JSON(http.StatusOK, findUsers)
}

func (u UsersController) FindByIdUser(c *gin.Context) {
	var findByIdUser models.FindByIdUser

	userIdU64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	userId := uint(userIdU64)
	findByIdUser.ID = userId

	if err := c.ShouldBindUri(&findByIdUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	err, user := services.FindByIdUser(findByIdUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u UsersController) CreateUsers(c *gin.Context) {
	var input models.CreateUser

	_ = c.BindJSON(&input)
	err, createUser := services.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, createUser)
}
