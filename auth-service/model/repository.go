package model

// Identity user identity
type Identity struct {
	Name    string `json:name`
	Email   string `json:email`
	Address string `json:address`
	Role    string `json:role`
}
