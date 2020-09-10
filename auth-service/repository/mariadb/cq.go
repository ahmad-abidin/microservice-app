package mariadb

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

// GetByUnP get identity by username and password
func (sql *repository) GetByUnP(username, password string) (*model.Claims, error) {
	stmt, err := sql.db.Prepare(`
		select i.name, i.email, i.address, r.name 
		from identity i, role r
		where i.email = ? and i.password = ? and i.id_role = r.id
	`)
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
		&i.Role,
	); err != nil {
		log.Printf("Error code R-GQ : %v", err)
		return nil, errors.New("R-GQ")
	}

	return i, nil
}
