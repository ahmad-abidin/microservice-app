package model

import (
	"errors"
	"log"
)

type Identity struct {
	Name    string `json:name`
	Email   string `json:email`
	Address string `json:address`
	Role    string `json:role`
}

func LogAndError(code string, err error) error {
	log.Printf("Error code %v : %v", code, err)
	return errors.New(code)
}
