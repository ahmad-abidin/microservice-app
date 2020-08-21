package model

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claims required to generate JWT
type Claims struct {
	jwt.StandardClaims
	Name    string
	Email   string
	Address string
}
