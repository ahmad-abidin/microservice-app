package usecase

import (
	"crypto/sha256"
	"fmt"

	"microservice-app/auth-service/model"
	"microservice-app/auth-service/repository/nosql"
	"microservice-app/auth-service/repository/sql"
	"microservice-app/auth-service/utils"
)

// Usecase ...
type Usecase interface {
	Authentication(string, string) (*string, error)
	Authorization(string) (*model.Identity, error)
}

type usecase struct {
	sqlRpo   sql.Repository
	nosqlRpo nosql.Repository
}

// NewUsecase ...
func NewUsecase(s sql.Repository, n nosql.Repository) Usecase {
	return &usecase{s, n}
}

// Authentication username (email)
func (u *usecase) Authentication(username, password string) (*string, error) {
	sha := sha256.New()
	sha.Write([]byte(password))
	password = fmt.Sprintf("%x", sha.Sum(nil))

	// get from redis
	identity, err := u.nosqlRpo.GetIdentity(username)
	if identity == nil {
		// get from mariadb
		identity, err = u.sqlRpo.GetIdentityByUnP(username, password)
		if err != nil {
			return nil, utils.WELI("e", "usecase-Aen_GIBUP", err)
		}

		// store to redis
		_, err := u.nosqlRpo.StoreIdentity(username, *identity)
		if err != nil {
			return nil, utils.WELI("e", "usecase-Aen_SI", err)
		}
	}

	encryptedIdentity, err := Encrypt(*identity)
	if err != nil {
		return nil, utils.WELI("e", "usecase-Aen_E", err)
	}

	return encryptedIdentity, nil
}

// Authorization ...
func (u *usecase) Authorization(encryptedIdentity string) (*model.Identity, error) {
	decryptedIdentity, err := Decrypt(encryptedIdentity)
	if err != nil {
		return nil, utils.WELI("e", "usecase-Aor_D", err)
	}

	return decryptedIdentity, nil
}
