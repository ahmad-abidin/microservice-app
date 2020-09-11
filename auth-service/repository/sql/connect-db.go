package sql

import (
	"database/sql"
	"fmt"
	"microservice-app/auth-service/model"
)

// ConnectDB ...
func ConnectDB(username, password, host, port, dbname, driver string) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open(driver, connString)
	if err != nil {
		return nil, model.Log("e", "sql-CDB_O", err)
	}

	if err := db.Ping(); err != nil {
		return nil, model.Log("e", "sql-CDB_P", err)
	}

	return db, nil
}
