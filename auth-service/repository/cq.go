package repository

import (
	"database/sql"
	"errors"
	"log"
	"microservice-app/auth-service/model"
)

type sqlRepository struct {
	db *sql.DB
}

// Repository ...
type Repository interface {
	GetByUnP(*model.Credential) (*model.Identity, error)
}

// NewRepository ...
func NewRepository(db *sql.DB) Repository {
	return &sqlRepository{db: db}
}

func (sql *sqlRepository) GetByUnP(c *model.Credential) (*model.Identity, error) {
	stmt, err := sql.db.Prepare("select name, email, address from identity where name = ? and password = ?")
	if err != nil {
		log.Fatalf("Error code GP : %v", err)
		return nil, errors.New("GP")
	}
	defer stmt.Close()

	i := new(model.Identity)
	if err := stmt.QueryRow(c.Username, c.Password).Scan(
		&i.Name,
		&i.Email,
		&i.Address,
	); err != nil {
		log.Fatalf("Error code GQ : %v", err)
		return nil, errors.New("GQ")
	}

	return i, nil
}
