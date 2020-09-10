package nosql

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"microservice-app/auth-service/model"

	"github.com/go-redis/redis"
)

type Repository interface {
	GetIdentity(string) (*model.Identity, error)
	StoreIdentity(string, model.Identity) (*string, error)
}

type repository struct {
	*redis.Client
}

// NewNoSQLRepository ...
func NewNoSQLRepository(client *redis.Client) Repository {
	return &repository{client}
}

func (r *repository) GetIdentity(key string) (*model.Identity, error) {
	stringIdentity, err := r.Get(key).Result()
	if fmt.Sprintf("%v", err) != "redis: nil" {
		return nil, err
	}
	if err != nil {
		log.Printf("Error code nosql-GI_G : %v", err)
		return nil, errors.New("nosql-GI_G")
	}

	identity := new(model.Identity)
	err = json.Unmarshal([]byte(stringIdentity), &identity)
	if err != nil {
		log.Printf("Error code nosql-GI_U : %v", err)
		return nil, errors.New("nosql-GI_U")
	}

	return identity, nil
}

func (r *repository) StoreIdentity(key string, identity model.Identity) (*string, error) {
	byteIdentity, err := json.Marshal(identity)
	if err != nil {
		log.Printf("Error code nosql-SI_M : %v", err)
		return nil, errors.New("nosql-SI_M")
	}

	res, err := r.Set(key, string(byteIdentity), 0).Result()
	if err != nil {
		log.Printf("Error code nosql-SI_S : %v", err)
		return nil, errors.New("nosql-SI_S")
	}

	return &res, nil
}
