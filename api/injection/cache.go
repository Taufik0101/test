package injection

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"test/api/utils"
)

func SetupRedisConnection() *redis.Client {
	redisHost := utils.EnvVar("REDIS_HOST", "")

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Username: utils.EnvVar("REDIS_USERNAME", ""),
		Password: utils.EnvVar("REDIS_PASSWORD", ""),
		DB:       0,
	})

	pong, errPong := client.Ping(context.TODO()).Result()
	fmt.Println(pong, errPong)

	return client
}
