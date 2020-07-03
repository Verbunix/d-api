package databases

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //Database

func init() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	//Build connection string
	var dbUri = fmt.Sprintf(
		"host=%s user=%s port=%s dbname=%s sslmode=disable password=%s",
		dbHost, dbUser, dbPort, dbName, dbPass,
	)
	fmt.Println("dbUri: \t", dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

//Returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
