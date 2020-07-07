package controllers

import (
	"net/http"

	"dating-api/databases"
	"dating-api/models"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func (u UsersController) FindUsers(c *gin.Context) {
	db := databases.GetDb()
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
