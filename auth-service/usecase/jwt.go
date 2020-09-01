package usecase

import (
	"errors"
	"log"

	"microservice-app/auth-service/model"

	jwt "github.com/dgrijalva/jwt-go"
)

// Encrypt jwt.go
func Encrypt(claims *model.Claims) (string, error) {
	token := jwt.NewWithClaims(
		model.JwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(model.JwtSecretKey)
	if err != nil {
		log.Fatalf("error when signed token: %v", err)
		return "", errors.New("internal server error")
	}

	return signedToken, nil
}

// Decrypt jwt.go
func Decrypt(unsignedToken string) (jwt.MapClaims, error) {
	signedToken, err := jwt.Parse(unsignedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("internal server error")
		} else if method != model.JwtSigningMethod {
			return nil, errors.New("internal server error")
		}
		return model.JwtSecretKey, nil
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

	return claims, nil
}
