package main

import (
	"d-api/migrations"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var router *gin.Engine

func main() {
	m := new(migrations.Migrations)
	m.Sync()

	router = gin.Default()

	initializeRoutes()

	port := os.Getenv("PORT")
	_ = router.Run(":" + port)
}
