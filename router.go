package main

import (
	"dating-api/controllers"
)

func initializeRoutes() {
	pingController := new(controllers.PingController)

	router.GET("/ping", pingController.Ping)
	userGroup := router.Group("users")
	{
		usersController := new(controllers.UsersController)
		userGroup.GET("", usersController.FindUsers)
	}
}
