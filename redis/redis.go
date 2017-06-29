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
	Port int
	DB   int
}

type Client *redis.Client

const Nil = redis.Nil

func NewClient(info *RedisConfig) Client {
	if info.Host == "" {
		panic(errors.New("redis config error"))
	}
	if info.Port == 0 {
		info.Port = 6379
	}
	var client *redis.Client
	for {
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", info.Host, info.Port),
			Password: "",
			DB:       0,
		})
		_, err := client.Ping().Result()
		if err != nil {
			loggers.Error.Printf("Failed to connect Redis Server: %s:%s", info.Host, info.Port)
			time.Sleep(2 * time.Second)
			loggers.Warn.Printf("Retrying to connect to redis")
		} else {
			break
		}
	}

	return client
}
