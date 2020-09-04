package repository

import (
	"database/sql"
	"errors"
	"log"
	"microservice-app/auth-service/model"
)

// Repository ...
type Repository interface {
	GetByUnP(model.Credential) (*model.Claims, error)
}

type repository struct {
	db *sql.DB
}

// NewRepository ...
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (sql *repository) GetByUnP(c model.Credential) (*model.Claims, error) {
	stmt, err := sql.db.Prepare("select name, email, address from identity where name = ? and password = ?")
	if err != nil {
		log.Fatalf("Error code R-GP : %v", err)
		return nil, errors.New("R-GP")
	}
	defer stmt.Close()

	i := new(model.Claims)
	if err := stmt.QueryRow(c.Username, c.Password).Scan(
		&i.Name,
		&i.Email,
		&i.Address,
	); err != nil {
		log.Fatalf("Error code R-GQ : %v", err)
		return nil, errors.New("R-GQ")
	}

	return i, nil
}
