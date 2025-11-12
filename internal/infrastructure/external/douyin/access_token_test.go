package douyin

import (
    "context"
    "testing"

    "github.com/joho/godotenv"
    "github.com/windorbird/xiang-backend/internal/infrastructure/persistence/redis"
    "github.com/windorbird/xiang-backend/log"
)
func init() {
    log.InitLogger()
    redis.InitRedis()
    err := godotenv.Load("../../../../.env")
    if err != nil {
        log.Logger.Fatal("Error loading .env file")
    }
}

func TestGetAccessToken(t *testing.T) {
    //err := godotenv.Load("../../../../.env")
    //if err != nil {
    //    log.Logger.Fatal("Error loading .env file")
    //}
    accessToken, err := GetAccessToken(context.Background())
    if err != nil {
        t.Errorf("GetAccessTokenFromDouyin() error = %v", err)
        return
    }
    if accessToken == "" {
        t.Errorf("GetAccessTokenFromDouyin() accessToken = %v", accessToken)
        return
    }
    t.Logf("GetAccessTokenFromDouyin() accessToken = %v", accessToken)
}

