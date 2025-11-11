package redis

var douYinPrefix = "douyin_"

func GetDouYinRedisKey(key string) string {
	return douYinPrefix + key
}
