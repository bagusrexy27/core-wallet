package config

import (
	"context"

	"github.com/core-wallet/app/utils"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     AppConfig.RedisHost + ":" + AppConfig.RedisPort,
		Password: "",
		DB:       0,
	})

	// Tes koneksi
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		return err
	}

	utils.InfoLog("INITIALIZE - APP Redis client initialized successfully")
	return nil
}
