package controllers

import (
	"d-api/models"
	"d-api/services/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func (u UsersController) FindUsers(c *gin.Context) {
	findUsers := users.FindAll()
	c.JSON(http.StatusOK, findUsers)
}

func (u UsersController) FindByIdUser(c *gin.Context) {
	var findByIdUser models.FindByIdUser

	userIdU64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	findByIdUser.ID = uint(userIdU64)

	if err := c.ShouldBindUri(&findByIdUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	err, user := users.FindById(findByIdUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u UsersController) CreateUsers(c *gin.Context) {
	var input models.CreateUser

	_ = c.BindJSON(&input)
	err, user := users.Create(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u UsersController) UpdateUser(c *gin.Context) {
	var input models.UpdateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err, user := users.Update(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u UsersController) DeleteUser(c *gin.Context) {
	var findByIdUser models.FindByIdUser

	userIdU64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	findByIdUser.ID = uint(userIdU64)

	if err := c.ShouldBindUri(&findByIdUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	err, _ = users.Delete(findByIdUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "User deleted successfully")
}
