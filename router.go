package main

import (
	"d-api/controllers"
	"d-api/middlewares"
)

func initializeRoutes() {
	auth := middlewares.JwtTokenAuthMiddleware()

	pingController := new(controllers.PingController)
	router.GET("/ping", pingController.Ping)

	userGroup := router.Group("users")
	{
		usersController := new(controllers.UsersController)
		userGroup.GET("", auth, usersController.FindUsers)
		userGroup.GET("/:id", auth, usersController.FindByIdUser)
		userGroup.POST("", auth, usersController.CreateUsers)
		userGroup.PATCH("", auth, usersController.UpdateUser)
		userGroup.DELETE("/:id", auth, usersController.DeleteUser)
	}

	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/signin", authController.SignIn)
}
