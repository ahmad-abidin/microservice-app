package repository

import (
	"database/sql"
	"errors"
	"log"
	"microservice-app/auth-service/model"
)

// Repository ...
type Repository interface {
	GetByUnP(string, string) (*model.Claims, error)
}

type repository struct {
	db *sql.DB
}

// NewRepository ...
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (sql *repository) GetByUnP(username, password string) (*model.Claims, error) {
	stmt, err := sql.db.Prepare("select name, email, address from identity where name = ? and password = ?")
	if err != nil {
		log.Printf("Error code R-GP : %v", err)
		return nil, errors.New("R-GP")
	}
	defer stmt.Close()

	i := new(model.Claims)
	if err := stmt.QueryRow(username, password).Scan(
		&i.Name,
		&i.Email,
		&i.Address,
	); err != nil {
		log.Printf("Error code R-GQ : %v", err)
		return nil, errors.New("R-GQ")
	}

	return i, nil
}
