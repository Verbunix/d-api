package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	initializeRoutes()

	port := os.Getenv("PORT")
	_ = router.Run(":" + port)
}
