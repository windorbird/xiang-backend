package redis_test

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	redis "github.com/windorbird/xiang-backend/internal/infrastructure/persistence/redis"
	"github.com/windorbird/xiang-backend/log"
    "go.uber.org/zap"

    "testing"
	"time"
)

func TestExpireTime(t *testing.T) {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		t.Errorf("load env file err: %v", err)
	}
	log.InitLogger()
	ctx := context.Background()
	// 初始化 Redis 客户端
	redis.InitRedis()

	redisKey := redis.GetDouYinRedisKey("bbb")
	err = redis.Client.Set(ctx, redisKey, "aabbccddeeff", 10*time.Minute).Err()
	if err != nil {
		t.Errorf("set access token to cache err: %v", err)
	}
	dd, err := redis.Client.ExpireTime(ctx, redisKey).Result()
	if err != nil {
		t.Errorf("set access token to cache err: %v", err)
	}
	fmt.Println(dd) //489664h6m34s 1970-01-01 00:00:00 距离现在的时间长度 xxhxxmxxs
	mm, ee := redis.Client.TTL(ctx, redisKey).Result()
	fmt.Println(ee)
	fmt.Println(mm > time.Duration(5*time.Minute))
	fmt.Println(mm > time.Duration(15*time.Minute))
}

func TestGetDouYinAccessToken(t *testing.T) {
    err := godotenv.Load("../../../../.env")
    if err != nil {
        t.Errorf("load env file err: %v", err)
    }
    log.InitLogger()
    ctx := context.Background()
    // 初始化 Redis 客户端
    redis.InitRedis()
    redisKey := redis.GetDouYinRedisKey("access_token")
    accessToken, err := redis.Client.Get(ctx, redisKey).Result()
    if err != nil {
        log.Logger.Warn("get douyin access token from cache err:", zap.Error(err))
        accessToken = "" // 空值情况下可以触发刷新
    }
    fmt.Println(accessToken)
}