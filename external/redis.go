package external

import (
	"github.com/redis/go-redis/v9"
)

// Redisに接続し，Clientオブジェクトを返す
func ConnectRedis() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return rdb, nil
}
