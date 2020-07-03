package controllers

import (
	"dating-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type UsersController struct{}

func (u UsersController) FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
