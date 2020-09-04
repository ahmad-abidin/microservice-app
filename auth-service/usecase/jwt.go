package usecase

import (
	"errors"
	"log"

	"microservice-app/auth-service/model"

	jwt "github.com/dgrijalva/jwt-go"
)

// Encrypt jwt.go
func Encrypt(c model.Claims) (*string, error) {
	c.StandardClaims.Issuer = model.AppicationName
	c.StandardClaims.ExpiresAt = model.ExpirationDuration

	token := jwt.NewWithClaims(
		model.JwtSigningMethod,
		c,
	)

	signedToken, err := token.SignedString(model.JwtSecretKey)
	if err != nil {
		log.Fatalf("Error code U-ES : %v", err)
		return nil, errors.New("U-ES")
	}

	return &signedToken, nil
}

// Decrypt jwt.go
func Decrypt(unsignedToken string) (*model.Claims, error) {
	signedToken, err := jwt.Parse(unsignedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Signing method invalid")
		} else if method != model.JwtSigningMethod {
			return nil, errors.New("Signing method invalid")
		}
		return model.JwtSecretKey, nil
	})

	if err != nil {
		log.Fatalf("Error code U-DP : %v", err)
		return nil, errors.New("U-DP")
	}

	claims, ok := signedToken.Claims.(jwt.MapClaims)
	if !ok || !signedToken.Valid {
		log.Fatalf("Error code U-DC : %v", ok)
		return nil, errors.New("U-DC")
	}

	c := new(model.Claims)
	for k, v := range claims {
		switch k {
		case "Name":
			c.Name = v.(string)
			break
		case "Email":
			c.Email = v.(string)
			break
		case "Address":
			c.Address = v.(string)
			break
		}
	}
	return c, nil
}
