package redis

import redisclient "github.com/MahmoudMekki/ChatSystem/clients/redis"

func CacheLastChatNumber(token string) (int, error) {
	redis := redisclient.GetRedisClient()
	last, err := redis.Incr(token).Result()
	return int(last), err
}
