package external

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

// Redisに接続し，Clientオブジェクトを返す
func ConnectRedis() (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return rdb, nil
}
