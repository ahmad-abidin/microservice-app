package usecase

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"

	"microservice-app/auth-service/model"
	"microservice-app/auth-service/repository/nosql"
	"microservice-app/auth-service/repository/sql"
)

type Usecase interface {
	Authentication(string, string) (*string, error)
	Authorization(string) (*model.Identity, error)
}

// Usecase ...
type usecase struct {
	sqlRpo   sql.Repository
	nosqlRpo nosql.Repository
}

func NewUsecase(s sql.Repository, n nosql.Repository) Usecase {
	return &usecase{s, n}
}

// Authentication username (email)
func (u *usecase) Authentication(username, password string) (*string, error) {
	sha := sha256.New()
	sha.Write([]byte(password))
	password = fmt.Sprintf("%x", sha.Sum(nil))

	// get from redis
	identity, _ := u.nosqlRpo.GetIdentity(username)

	if identity == nil {
		// get from mariadb
		identity, err := u.sqlRpo.GetIdentityByUnP(username, password)
		if err != nil {
			log.Printf("Error code usecase-Aen_GIBUP <- %v", err)
			return nil, errors.New("usecase-Aen_GIBUP")
		}

		// store to redis
		status, err := u.nosqlRpo.StoreIdentity(username, *identity)
		if err != nil {
			log.Printf("Error code usecase-Aen_SI <- %v", err)
			return nil, errors.New("usecase-Aen_SI")
		}

		log.Printf("Successfully caching data : %v", *status)
	}

	encryptedIdentity, err := Encrypt(*identity)
	if err != nil {
		log.Printf("Error code usecase-Aen_E <- %v", err)
		return nil, errors.New("usecase-Aen_E")
	}

	return encryptedIdentity, nil
}

// Authorization ...
func (u *usecase) Authorization(encryptedIdentity string) (*model.Identity, error) {
	decryptedIdentity, err := Decrypt(encryptedIdentity)
	if err != nil {
		log.Printf("Error code usecase-Aor_D <- %v", err)
		return nil, errors.New("usecase-Aor_D")
	}

	return decryptedIdentity, nil
}
