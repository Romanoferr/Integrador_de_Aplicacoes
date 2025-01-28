package connectdb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() (*sql.DB, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "gomysql",
		Passwd: "gomysql",
		Addr:   "127.0.0.1:3306",
		DBName: "statusinvest",
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
