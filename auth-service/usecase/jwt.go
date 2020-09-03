package usecase

import (
	"errors"
	"log"

	"microservice-app/auth-service/model"

	jwt "github.com/dgrijalva/jwt-go"
)

// Encrypt jwt.go
func Encrypt(i *model.Identity) (string, error) {
	c := new(model.Claims)
	c.Name = i.Name
	c.Email = i.Email
	c.Address = i.Address
	c.StandardClaims.Issuer = model.AppicationName
	c.StandardClaims.ExpiresAt = model.ExpirationDuration

	token := jwt.NewWithClaims(
		model.JwtSigningMethod,
		c,
	)

	signedToken, err := token.SignedString(model.JwtSecretKey)
	if err != nil {
		log.Fatalf("Error code ES : %v", err)
		return "", errors.New("ES")
	}

	return signedToken, nil
}

// Decrypt jwt.go
func Decrypt(unsignedToken string) (jwt.MapClaims, error) {
	signedToken, err := jwt.Parse(unsignedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Signing method invalid")
		} else if method != model.JwtSigningMethod {
			return nil, errors.New("Signing method invalid")
		}
		return model.JwtSecretKey, nil
	})

	if err != nil {
		log.Fatalf("Error code DP : %v", err)
		return nil, errors.New("DP")
	}

	claims, ok := signedToken.Claims.(jwt.MapClaims)
	if !ok || !signedToken.Valid {
		log.Fatalf("Error code DC : %v", err)
		return nil, errors.New("DC")
	}

	return claims, nil
}
