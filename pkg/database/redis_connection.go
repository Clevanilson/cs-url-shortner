package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisConnection struct {
	client *redis.Client
}

func NewRedisConnection() *RedisConnection {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	client.Set(ctx, "global:id", 999, 0)
	return &RedisConnection{client}
}

func (c *RedisConnection) Set(key, value string) error {
	if err := c.client.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (c *RedisConnection) Get(key string) (string, error) {
	result, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", nil
	}
	return result, nil
}

func (c *RedisConnection) Increment(key string) (int64, error) {
	result, err := c.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *RedisConnection) Remove(key string) error {
	if err := c.client.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
