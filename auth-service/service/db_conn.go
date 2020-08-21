package service

import (
	"database/sql"
	"log"

	// mysql/mariaDB driver
	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB ...
func ConnectDB() (*sql.DB, error) {
	log.Println("connecting to database...")

	db, err := sql.Open("mysql", "root:root@tcp(db_user:3306)/user")
	if err != nil {
		log.Fatalf("failed to open connection to db_user: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping to db_user: %v", err)
		return nil, err
	}

	return db, nil
}
