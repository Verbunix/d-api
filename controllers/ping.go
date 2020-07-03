package controllers

import "github.com/gin-gonic/gin"

type PingController struct{}

func (p PingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
