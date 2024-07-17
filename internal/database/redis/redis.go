package redis

import (
	"context"
	"fmt"
	"github.com/mostafababaii/gorest/config"
	"github.com/redis/go-redis/v9"
)

var DefaultClient *redis.Client

func Setup() {
	redisAddress := fmt.Sprintf("%s:%d", config.RedisConfig.Host, config.RedisConfig.Port)

	redisOptions := redis.Options{
		Addr:            redisAddress,
		DB:              0,
		Username:        config.RedisConfig.User,
		Password:        config.RedisConfig.Password,
		MaxIdleConns:    config.RedisConfig.MaxIdle,
		MaxActiveConns:  config.RedisConfig.MaxIdle,
		ConnMaxIdleTime: config.RedisConfig.IdleTimeout,
	}

	DefaultClient = redis.NewClient(&redisOptions)
	fmt.Println(DefaultClient.Ping(context.Background()))
}
