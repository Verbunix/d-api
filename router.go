package main

import (
	"dating-api/controllers"
)

func initializeRoutes() {
	pingController := new(controllers.PingController)

	router.GET("/ping", pingController.Ping)
}
