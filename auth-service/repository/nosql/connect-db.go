package nosql

import (
	"microservice-app/auth-service/model"

	"github.com/go-redis/redis"
)

// ConnectDB ...
func ConnectDB(host, port, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	if _, err := rdb.Ping().Result(); err != nil {
		return nil, model.Log("e", "nosql-CDB_NC", err)
	}

	return rdb, nil
}
