package service

import (
	"context"
	"errors"
	"log"
	"time"

	"microservice-app/auth-service/model"

	jwt "github.com/dgrijalva/jwt-go"

	"google.golang.org/grpc/metadata"
)

// Server ...
type Server struct{}

// Authentication ...
func (s *Server) Authentication(ctx context.Context, identity *model.Identity) (*model.Credential, error) {
	log.Println("### Authentication ###")

	claims := model.Claims{}

	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("error Authentication-ConnectDB: %v", err)
		return nil, errors.New("internal server error")
	}
	defer db.Close()

	log.Println("Authenticating...")

	err = db.QueryRow(
		`select name, email, address from identity where name = ? and password = ?`,
		identity.Username,
		identity.Password).Scan(
		&claims.Name,
		&claims.Email,
		&claims.Address)
	if err != nil {
		log.Fatalf("error Authentication-QueryRow: %v", err)
		return nil, errors.New("invalid username and password")
	}

	log.Println("Createing Token...")

	claims.StandardClaims.Issuer = "simple-microservice-app"
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(1) * time.Hour).Unix()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, err := token.SignedString([]byte("namanya juga secret key, bebas"))
	if err != nil {
		log.Fatalf("error when signed token: %v", err)
		return nil, errors.New("internal server error")
	}

	log.Println("### Succesfully Authentication ###")

	return &model.Credential{
		Token: signedToken,
	}, nil
}

// Authorization ...
func (s *Server) Authorization(ctx context.Context, credential *model.Credential) (*model.FullIdentity, error) {
	log.Println("### Authorization ###")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalf("error geting metadata")
		return nil, errors.New("internal server error")
	}

	log.Println("Authorizating...")
	arrayOfMd := md.Get("authorization")
	unsignedToken := arrayOfMd[0]
	signedToken, err := jwt.Parse(unsignedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("internal server error")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("internal server error")
		}
		return []byte("namanya juga secret key, bebas"), nil
	})
	if err != nil {
		log.Fatalf("error when signing token: %v", err)
		return nil, errors.New("internal server error")
	}
	claims, ok := signedToken.Claims.(jwt.MapClaims)
	if !ok || !signedToken.Valid {
		log.Fatalf("error when check claims: %v", err)
		return nil, errors.New("internal server error")
	}
	fi := model.FullIdentity{}
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
