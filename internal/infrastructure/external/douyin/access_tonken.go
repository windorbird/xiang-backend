package douyin

import (
    "context"
    "errors"
    "os"
    "time"

    credential "github.com/bytedance/douyin-openapi-credential-go/client"
    openApiSdkClient "github.com/bytedance/douyin-openapi-sdk-go/client"
    redis "github.com/windorbird/xiang-backend/internal/infrastructure/persistence/redis"
    "github.com/windorbird/xiang-backend/log"
    "go.uber.org/zap"
)

func GetAccessToken(ctx context.Context) (string, error) {
	accessToken, err := GetAccessTokenFromCache(ctx)
	if err == nil && accessToken != "" {
		return accessToken, nil
	}
	accessToken, err = GetAccessTokenFromDouyin() //todo: 并发问题
	if err != nil {
		return "", err
	}

	if accessToken == "" {
		log.Logger.Error("sdk call err:", zap.Error(errors.New("douyin access token is empty")))
		return "", errors.New("douyin access_token is empty")
	}
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Logger.Error("set douyin access token to cache panic:", zap.Any("recover", r))
            }
        }()
        err := SetAccessTokenToCache(ctx, accessToken)
        if err != nil {
            log.Logger.Warn("set douyin access token to cache err:", zap.Error(err))
        }
    }()
	return accessToken, nil
}

func SetAccessTokenToCache(ctx context.Context, accessToken string) error {
    expiration := time.Hour * 2
    redisKey := redis.GetDouYinRedisKey("access_token")
    return redis.Client.Set(ctx, redisKey, accessToken, expiration).Err() //todo:连接池
}

func GetAccessTokenFromCache(ctx context.Context) (string, error) {
    redisKey := redis.GetDouYinRedisKey("access_token")
    accessToken, err := redis.Client.Get(ctx, redisKey).Result() //todo:redis连接池
    if err != nil {
        log.Logger.Warn("get douyin access token from cache err:", zap.Error(err))
    }
    return accessToken, err
}

// GetAccessTokenFromDouyin 通过抖音SDK获取 accessToken
func GetAccessTokenFromDouyin() (string, error) {
    appId := os.Getenv("DouYinAppID")
    appSecret := os.Getenv("DouYinAppSecret")
    opt := new(credential.Config).
        SetClientKey(appId).
        SetClientSecret(appSecret)

    sdkClient, err := openApiSdkClient.NewClient(opt)
    if err != nil {
        log.Logger.Error("sdk init err:", zap.Error(err))
        return "", err
    }

    req := &openApiSdkClient.V2TokenRequest{}
    req.SetAppid(appId)
    req.SetGrantType("client_credential")
    req.SetSecret(appSecret)

    resp, err := sdkClient.V2Token(req)
    if err != nil {
        log.Logger.Error("sdk call err:", zap.Error(err))
        return "", err
    }

    if *resp.ErrNo != 0 {
        err := errors.New(*resp.ErrTips)
        log.Logger.Error("sdk ErrTips:", zap.Error(err))
        return "", err
    }
    return *resp.Data.AccessToken, nil
}
