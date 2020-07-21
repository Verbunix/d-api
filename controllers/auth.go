package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{}

func (u UsersController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (u UsersController) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
