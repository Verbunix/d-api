package migrations

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	migrate "github.com/rubenv/sql-migrate"
)

type Migrations struct{}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Database struct {
	host, user, password, dbname, ssl string
	port                              int
	db                                *sql.DB
}

func (d *Database) Open() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		d.host,
		d.port,
		d.user,
		d.password,
		d.dbname,
	)
	if d.ssl == "disable" {
		psqlInfo = psqlInfo + " " + "sslmode=disable"
	}
	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)
	d.db = db
}

func (d *Database) IsOpened() {
	checkError(d.db.Ping())
}

func connection() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbSslmode := os.Getenv("DB_SSLMODE")
	database := Database{host: dbHost, port: dbPort, user: dbUser, password: dbPass, dbname: dbName, ssl: dbSslmode}
	database.Open()
	database.IsOpened()
	return database.db
}

func (m *Migrations) Sync() {
	migrate.SetTable("migrations")

	//Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	db := connection()

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
		fmt.Println(err)
	}

	_ = db.Close()

	fmt.Printf("Applied %d migrations!\n", n)
}
