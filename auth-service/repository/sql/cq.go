package sql

import (
	"database/sql"
	"microservice-app/auth-service/model"
)

// Repository ...
type Repository interface {
	GetIdentityByUnP(string, string) (*model.Identity, error)
}

type repository struct {
	*sql.DB
}

// NewSQLRepository ...
func NewSQLRepository(db *sql.DB) Repository {
	return &repository{db}
}

// GetByUnP get identity by username and password
func (r *repository) GetIdentityByUnP(username, password string) (*model.Identity, error) {
	stmt, err := r.Prepare(`
		select i.name, i.email, i.address, r.name 
		from identity i, role r
		where i.email = ? and i.password = ? and i.id_role = r.id
	`)
	if err != nil {
		return nil, model.Log("e", "sql-GIBUP_P", err)
	}
	defer stmt.Close()

	identity := new(model.Identity)
	if err := stmt.QueryRow(username, password).Scan(
		&identity.Name,
		&identity.Email,
		&identity.Address,
		&identity.Role,
	); err != nil {
		return nil, model.Log("e", "sql-GIBUP_QR", err)
	}

	return identity, nil
}
