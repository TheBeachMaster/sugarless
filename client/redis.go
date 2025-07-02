package client

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/redis/go-redis/v9"
)

func NewCacheDBClient(address string) (*redis.Client, error) {

	redisOpts := &redis.Options{
		Addr:         address,
		MinIdleConns: 1,
		MaxIdleConns: runtime.NumCPU(),
	}

	redisClient := redis.NewClient(redisOpts)
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Printf("unable to connect to cache %s", err.Error())
		return nil, fmt.Errorf("could not reach cache server ")
	}

	return redisClient, nil
}
