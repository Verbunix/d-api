package main

import "dating-api/controllers"

func initializeRoutes() {
	pingController := new(controllers.PingController)
	router.GET("/ping", pingController.Ping)

	userGroup := router.Group("users")
	{
		usersController := new(controllers.UsersController)
		userGroup.GET("", usersController.FindUsers)
		userGroup.GET("/:id", usersController.FindByIdUser)
		userGroup.POST("", usersController.CreateUsers)
		userGroup.PATCH("", usersController.UpdateUser)
		userGroup.DELETE("/:id", usersController.DeleteUser)
	}

	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/signin", authController.SignIn)
}
