package usecase

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"microservice-app/auth-service/model"
	"microservice-app/auth-service/repository/nosql"
	"microservice-app/auth-service/repository/sql"
)

// Usecase ...
type Usecase interface {
	Authentication(string) (*string, error)
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
func (u *usecase) Authentication(basicAuth string) (*string, error) {
	// decrypt basic auth
	basicAuth = strings.Replace(basicAuth, "Basic ", "", -1)
	decodedByteBasicAuth, err := base64.StdEncoding.DecodeString(basicAuth)
	if decodedByteBasicAuth == nil {
		return nil, model.Log("e", "usecase-Aen_DS", err)
	}
	decodedStringBasicAuth := string(decodedByteBasicAuth)
	i := strings.Index(decodedStringBasicAuth, ":")
	if i == -1 {
		return nil, model.Log("e", "usecase-Aen_I", err)
	}
	username, password := decodedStringBasicAuth[0:i], decodedStringBasicAuth[i+1:]

	// hash pasword
	sha := sha256.New()
	sha.Write([]byte(password))
	password = fmt.Sprintf("%x", sha.Sum(nil))

	// get from redis
	identity, err := u.nosqlRpo.GetIdentity(username)
	if identity == nil {
		// get from mariadb
		identity, err = u.sqlRpo.GetIdentityByUnP(username, password)
		if err != nil {
			return nil, model.Log("e", "usecase-Aen_GIBUP", err)
		}

		// store to redis
		_, err := u.nosqlRpo.StoreIdentity(username, *identity)
		if err != nil {
			return nil, model.Log("e", "usecase-Aen_SI", err)
		}

		model.Log("s", "usecase-Aen_SI", fmt.Errorf("Successfully cached to redis"))
	}

	encryptedIdentity, err := Encrypt(*identity)
	if err != nil {
		return nil, model.Log("e", "usecase-Aen_E", err)
	}

	return encryptedIdentity, nil
}

// Authorization ...
func (u *usecase) Authorization(encryptedIdentity string) (*model.Identity, error) {
	encryptedIdentity = strings.Replace(encryptedIdentity, "Bearer ", "", -1)

	decryptedIdentity, err := Decrypt(encryptedIdentity)
	if err != nil {
		return nil, model.Log("e", "usecase-Aor_D", err)
	}

	return decryptedIdentity, nil
}
