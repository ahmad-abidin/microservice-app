package usecase

import (
	"context"
	"errors"
	"log"

	"microservice-app/auth-service/model"
	repo "microservice-app/auth-service/repository"

	"google.golang.org/grpc/metadata"
)

// Usecase ...
type usecase struct {
	repository repo.Repository
}

func NewUsecase(r repo.Repository) usecase {
	return usecase{repository: r}
}

// Authentication ...
func (u *usecase) Authentication(ctx context.Context, c *model.Credential) (*model.Token, error) {
	t := new(model.Token)

	i, err := u.repository.GetByUnP(c)
	if err != nil {
		log.Fatalf("Error code AenG <- %v", err)
		return nil, errors.New("AenG")
	}

	signedToken, err := Encrypt(i)
	if err != nil {
		log.Fatalf("Error code AenE <- %v", err)
		return nil, errors.New("AenE")
	}

	t.Jwt = signedToken

	log.Printf("### Succesfully Authentication ###")
	return t, nil
}

// Authorization ...
func (u *usecase) Authorization(ctx context.Context, t *model.Token) (*model.Identity, error) {
	i := new(model.Identity)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalf("Error code AorF <- %v", ok)
		return nil, errors.New("AorF")
	}
	arrayOfMd := md.Get("authorization")
	unsignedToken := arrayOfMd[0]

	claims, err := Decrypt(unsignedToken)
	if err != nil {
		log.Fatalf("Error code AorD <- %v", err)
		return nil, errors.New("AorD")
	}

	for key, val := range claims {
		switch key {
		case "Name":
			i.Name = val.(string)
			break
		case "Email":
			i.Email = val.(string)
			break
		case "Address":
			i.Address = val.(string)
			break
		}
	}

	log.Printf("### Succesfully Authorization ###")
	return i, nil
}
