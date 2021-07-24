package redisCache

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func (r *Redis) Connect() {
	r.client = redis.NewClient(&redis.Options{
		Addr:     ADDRESS,
		Password: "",
		DB:       0,
	})
}

// func (r *Redis) Connect(address string) {
// 	r.client = redis.NewClient(&redis.Options{
// 		Addr:     address,
// 		Password: "",
// 		DB:       0,
// 	})
// }

func (r *Redis) GetData(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		fmt.Println("Remote cache miss")
		return "", fmt.Errorf("remote cache miss")
	}
	return val, nil
}

func (r *Redis) SetData(key string, value string) error {
	return r.client.Set(key, value, 0).Err()
}

func (r *Redis) DeleteData(key string) error {
	return r.client.Del(key).Err()
}
