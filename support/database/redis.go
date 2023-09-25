package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	CreateClient()
	GetClient() *redis.Client
	GetContext() context.Context
	DisconnectClient()
}

type DefaultRedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient() *DefaultRedisClient {
	return &DefaultRedisClient{
		ctx: context.Background(),
	}
}

func (c *DefaultRedisClient) CreateClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	c.client = rdb
}

func (c DefaultRedisClient) GetClient() *redis.Client {
	return c.client
}

func (c *DefaultRedisClient) GetContext() context.Context {
	return c.ctx
}

func (c *DefaultRedisClient) DisconnectClient() {
	c.client.Close()
}
