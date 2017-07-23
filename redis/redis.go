package redis

import (
	"errors"
	"fmt"
	"fxlibraries/loggers"
	"gopkg.in/redis.v5"
	"time"
)

type RedisConfig struct {
	Host string
	Port string
	DB   int
}

type Client *redis.Client

const Nil = redis.Nil

const RetryCount = 5

func NewClient(info *RedisConfig) Client {
	if info.Host == "" {
		panic(errors.New("redis config error"))
	}
	if info.Port == "" {
		info.Port = "6379"
	}
	var (
		client *redis.Client
		err    error
	)
	for i := 0; i < RetryCount; i++ {
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", info.Host, info.Port),
			Password: "",
			DB:       info.DB,
		})
		_, err = client.Ping().Result()
		if err != nil {
			loggers.Error.Printf("Failed to connect Redis Server: %v", info)
			time.Sleep(2 * time.Second)
			loggers.Warn.Printf("Retrying to connect to redis: %v", info)
		} else {
			return client
		}
	}
	panic(err)
}
