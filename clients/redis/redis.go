package redisclient

import (
	"fmt"
	"github.com/MahmoudMekki/ChatSystem/config"
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

var redisClient *redis.Client

func connectToRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetEnvVar("REDIS_URL"),
		Password: config.GetEnvVar("REDIS_PASSWORD"),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg(fmt.Sprintf("Redis is running on -> %s", config.GetEnvVar("REDIS_URL")))
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		connectToRedis()
	}
	return redisClient
}
