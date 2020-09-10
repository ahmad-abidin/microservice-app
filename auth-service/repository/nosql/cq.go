package nosql

import (
	"encoding/json"
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
	if err != nil {
		return nil, model.LogAndError("nosql-GI_G", err)
	}

	identity := new(model.Identity)
	err = json.Unmarshal([]byte(stringIdentity), &identity)
	if err != nil {
		return nil, model.LogAndError("nosql-GI_U", err)
	}

	return identity, nil
}

func (r *repository) StoreIdentity(key string, identity model.Identity) (*string, error) {
	byteIdentity, err := json.Marshal(identity)
	if err != nil {
		return nil, model.LogAndError("nosql-SI_M", err)
	}

	res, err := r.Set(key, string(byteIdentity), 0).Result()
	if err != nil {
		return nil, model.LogAndError("nosql-SI_S", err)
	}

	return &res, nil
}
