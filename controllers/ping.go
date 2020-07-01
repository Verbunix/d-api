package controllers

import "github.com/gin-gonic/gin"

type PingController struct{}

func (p PingController) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
