package redis

import (
	"fmt"
	redisclient "github.com/MahmoudMekki/ChatSystem/clients/redis"
)

func CacheLastChatNumber(token string) (int, error) {
	redis := redisclient.GetRedisClient()
	last, err := redis.Incr(token).Result()
	return int(last), err
}

func CacheLastMsgNumber(token string, chatNum int) (int, error) {
	redis := redisclient.GetRedisClient()
	last, err := redis.Incr(generateMsgKey(token, chatNum)).Result()
	return int(last), err
}

func generateMsgKey(token string, chatNum int) string {
	return fmt.Sprintf("%s%d", token, chatNum)
}
