package usecase

import (
	"errors"
	"log"

	"microservice-app/auth-service/model"

	jwt "github.com/dgrijalva/jwt-go"
)

// Encrypt jwt.go
func Encrypt(identity model.Identity) (*string, error) {
	claims := model.Claims{
		jwt.StandardClaims{
			Issuer:    model.AppicationName,
			ExpiresAt: model.ExpiredDuration,
		},
		identity,
	}

	token := jwt.NewWithClaims(
		model.JwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(model.JwtSecretKey)
	if err != nil {
		log.Printf("Error code usecase-E_SS : %v", err)
		return nil, errors.New("usecase-E_SS")
	}

	return &signedToken, nil
}

// Decrypt jwt.go
func Decrypt(unsignedToken string) (*model.Identity, error) {
	signedToken, err := jwt.Parse(unsignedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Signing method invalid")
		} else if method != model.JwtSigningMethod {
			return nil, errors.New("Signing method invalid")
		}
		return model.JwtSecretKey, nil
	})
	if err != nil {
		log.Printf("Error code usecase-D_P : %v", err)
		return nil, errors.New("usecase-D_P")
	}

	claims, ok := signedToken.Claims.(jwt.MapClaims)
	if !ok || !signedToken.Valid {
		log.Printf("Error code usecase-D_C : %v", ok)
		return nil, errors.New("usecase-D_C")
	}

	i := new(model.Identity)
	for k, v := range claims {
		switch k {
		case "Name":
			i.Name = v.(string)
			break
		case "Email":
			i.Email = v.(string)
			break
		case "Address":
			i.Address = v.(string)
			break
		case "Role":
			i.Role = v.(string)
			break
		}
	}
	return i, nil
}
