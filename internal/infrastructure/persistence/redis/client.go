package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

    "github.com/redis/go-redis/v9/maintnotifications"
    "github.com/windorbird/xiang-backend/utils"

	"github.com/redis/go-redis/v9"
	"github.com/windorbird/xiang-backend/log"
	"go.uber.org/zap"
)

var Client *redis.Client

type Config struct {
	Addr     string // 地址：host:port
	Username string // 用户名（6.0+）
	Password string // 密码
	DB       int    // 数据库编号
}

// InitRedis 初始化Redis客户端
func InitRedis() error {
	cfg, err := loadRedisConfig()
	if err != nil {
		log.Logger.Fatal("加载Redis配置失败", zap.Error(err))
        return err
	}
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
        MaintNotificationsConfig: &maintnotifications.Config{Mode: maintnotifications.ModeDisabled},
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

    _, err = Client.Ping(ctx).Result()
	if err != nil {
		log.Logger.Fatal("Redis连接失败", zap.Error(err))
        return err
	}
	log.Logger.Info("Redis连接成功", zap.String("addr", cfg.Addr), zap.Int("db", cfg.DB))
    return nil
}

// LoadRedisConfig 从环境变量中加载并解析Redis配置
func loadRedisConfig() (*Config, error) {
	db, err := strconv.Atoi(utils.GetEnv("REDIS_DB", "0")) // 默认为0
	if err != nil {
		return nil, fmt.Errorf("REDIS_DB格式错误: %v", err)
	}

	return &Config{
		Addr:     utils.GetEnv("REDIS_ADDR", "127.0.0.1:6379"), // 默认本地地址
		Username: utils.GetEnv("REDIS_USERNAME", "an"),         // 默认为空（低版本兼容）
		Password: utils.GetEnv("REDIS_PASSWORD", ""),           // 密码必填，建议在.env中设置
		DB:       db,
	}, nil
}
