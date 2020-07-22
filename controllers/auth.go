package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{}

func (u AuthController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (u AuthController) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
