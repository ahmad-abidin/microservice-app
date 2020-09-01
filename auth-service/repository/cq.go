package repository

import (
	"database/sql"
	"microservice-app/auth-service/model"
)

type sqlRepository struct {
	db *sql.DB
}

// Repository ...
type Repository interface {
	GetByUnP(*model.Identity) (*model.Claims, error)
}

// NewRepository ...
func NewRepository(db *sql.DB) Repository {
	return &sqlRepository{db: db}
}

func (sql *sqlRepository) GetByUnP(i *model.Identity) (*model.Claims, error) {
	stmt, err := sql.db.Prepare("select name, email, address from identity where name = ? and password = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	c := new(model.Claims)
	if err := stmt.QueryRow(i.Username, i.Password).Scan(
		&c.Name,
		&c.Email,
		&c.Address,
	); err != nil {
		return nil, err
	}

	return c, nil
}
