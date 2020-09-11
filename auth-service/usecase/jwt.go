package usecase

import (
	"fmt"
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
		return nil, model.Log("e", "usecase-E_SS", err)
	}

	return &signedToken, nil
}

// Decrypt jwt.go
func Decrypt(unsignedToken string) (*model.Identity, error) {
	signedToken, err := jwt.Parse(unsignedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.Log("e", "usecase-D_P", fmt.Errorf("signingmethod %v", ok))
		} else if method != model.JwtSigningMethod {
			return nil, model.Log("e", "usecase-D_P", fmt.Errorf("signingmethod %v", ok))
		}
		return model.JwtSecretKey, nil
	})
	if err != nil {
		return nil, model.Log("e", "usecase-D_P", err)
	}

	claims, ok := signedToken.Claims.(jwt.MapClaims)
	if !ok || !signedToken.Valid {
		return nil, model.Log("e", "usecase-D_C", err)
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
