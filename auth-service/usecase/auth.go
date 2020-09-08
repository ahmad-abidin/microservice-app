package usecase

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"

	"microservice-app/auth-service/model"
	rpo "microservice-app/auth-service/repository/sql"
)

type Usecase interface {
	Authentication(string, string) (*string, error)
	Authorization(string) (*model.Claims, error)
}

// Usecase ...
type usecase struct {
	repository rpo.Repository
}

func NewUsecase(r rpo.Repository) Usecase {
	return &usecase{r}
}

// Authentication ...
func (u *usecase) Authentication(username, password string) (*string, error) {
	sha := sha256.New()
	sha.Write([]byte(password))
	password = fmt.Sprintf("%x", sha.Sum(nil))

	claims, err := u.repository.GetByUnP(username, password)
	if err != nil {
		log.Printf("Error code U-AenG <- %v", err)
		return nil, errors.New("U-AenG")
	}

	signedToken, err := Encrypt(*claims)
	if err != nil {
		log.Printf("Error code U-AenE <- %v", err)
		return nil, errors.New("U-AenE")
	}

	return signedToken, nil
}

// Authorization ...
func (u *usecase) Authorization(t string) (*model.Claims, error) {
	claims, err := Decrypt(t)
	if err != nil {
		log.Printf("Error code U-AorD <- %v", err)
		return nil, errors.New("U-AorD")
	}

	return claims, nil
}
