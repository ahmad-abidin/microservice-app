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
type Usecase struct {
	Repository repo.Repository
}

// Authentication ...
func (u *Usecase) Authentication(ctx context.Context, identity *model.Identity) (*model.Credential, error) {
	log.Println("### Authentication ###")

	claims, err := u.Repository.GetByUnP(identity)
	if err != nil {
		log.Fatalf("error Authentication-GetByUnP: %v", err)
		return nil, errors.New("internal server error")
	}
	claims.StandardClaims.Issuer = model.AppicationName
	claims.StandardClaims.ExpiresAt = model.ExpirationDuration

	log.Println("Authenticating...")

	signedToken, err := Encrypt(claims)

	log.Println("### Succesfully Authentication ###")

	return &model.Credential{
		Token: signedToken,
	}, nil
}

// Authorization ...
func (u *Usecase) Authorization(ctx context.Context, credential *model.Credential) (*model.FullIdentity, error) {
	log.Println("### Authorization ###")

	fi := model.FullIdentity{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalf("error geting metadata")
		return nil, errors.New("internal server error")
	}
	arrayOfMd := md.Get("authorization")
	unsignedToken := arrayOfMd[0]

	log.Println("Authorizating...")

	claims, err := Decrypt(unsignedToken)
	if err != nil {
		log.Fatalf("error Decode token: %v", err)
		return nil, errors.New("invalid username and password")
	}

	for key, val := range claims {
		switch key {
		case "Name":
			fi.Name = val.(string)
			break
		case "Email":
			fi.Email = val.(string)
			break
		case "Address":
			fi.Address = val.(string)
			break
		}
	}

	log.Printf("### Succesfully Authorization ###")

	return &fi, nil
}
