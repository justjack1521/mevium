package mevcon

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	DSN      string `required:"true"`
	Password string `required:"true"`
}

func NewRedisConnection(ctx context.Context, config RedisConfig) (*redis.Client, error) {
	dsn := config.DSN
	if len(dsn) == 0 {
		return nil, errors.New("invalid redis dsn")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: config.Password,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
