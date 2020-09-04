package usecase

import (
	"errors"
	"log"

	"microservice-app/auth-service/model"
	repo "microservice-app/auth-service/repository"
)

type Usecase interface {
	Authentication(model.Credential) (*string, error)
	Authorization(string) (*model.Claims, error)
}

// Usecase ...
type usecase struct {
	repository repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{r}
}

// Authentication ...
func (u *usecase) Authentication(c model.Credential) (*string, error) {
	claims, err := u.repository.GetByUnP(c)
	if err != nil {
		log.Fatalf("Error code U-AenG <- %v", err)
		return nil, errors.New("U-AenG")
	}

	signedToken, err := Encrypt(*claims)
	if err != nil {
		log.Fatalf("Error code U-AenE <- %v", err)
		return nil, errors.New("U-AenE")
	}

	return signedToken, nil
}

// Authorization ...
func (u *usecase) Authorization(t string) (*model.Claims, error) {
	claims, err := Decrypt(t)
	if err != nil {
		log.Fatalf("Error code U-AorD <- %v", err)
		return nil, errors.New("U-AorD")
	}

	return claims, nil
}
