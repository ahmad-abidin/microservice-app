package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	JwtSigningMethod   = jwt.SigningMethodHS256
	JwtSecretKey       = []byte("uvuvwevwevwe onyetenyevwe ugwemuhwem osas")
	AppicationName     = "simple-microservice-app"
	ExpirationDuration = time.Now().Add(time.Duration(1) * time.Hour).Unix()
)

// Claims required to generate JWT
type Claims struct {
	jwt.StandardClaims
	// Identity
	Name    string
	Email   string
	Address string
	Role    string
}
