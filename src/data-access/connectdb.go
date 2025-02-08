package connectdb

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func ConnectDB() (*sql.DB, error) {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DOCKER_DB_USER"),
		Passwd: os.Getenv("DOCKER_DB_PASSWORD"),
		Addr:  	os.Getenv("DB_ADDR"),
		DBName: os.Getenv("DOCKER_DB_NAME"),
	}
	// Get a database handle.
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return Db, nil;

}
