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

func (r RedisConfig) NewRedisConnection(ctx context.Context) (*redis.Client, error) {
	dsn := r.DSN
	if len(dsn) == 0 {
		return nil, errors.New("invalid redis dsn")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: r.Password,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
