package redis

import (
	"github.com/go-redis/redis"
)

func ConnectDB(host, port, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	if _, err := rdb.Ping(); err != nil {
		return nil, err
	}

	return rdb, nil
}
